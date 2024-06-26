package keeper

import (
	"context"
	"fmt"

	math "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pokt-network/smt"

	"github.com/pokt-network/poktroll/telemetry"
	apptypes "github.com/pokt-network/poktroll/x/application/types"
	prooftypes "github.com/pokt-network/poktroll/x/proof/types"
	suppliertypes "github.com/pokt-network/poktroll/x/supplier/types"
	"github.com/pokt-network/poktroll/x/tokenomics/types"
)

// SettleSessionAccounting is responsible for all of the post-session accounting
// necessary to burn, mint or transfer tokens depending on the amount of work
// done. The amount of "work done" complete is dictated by `sum` of `root`.
//
// ASSUMPTION: It is assumed the caller of this function validated the claim
// against a proof BEFORE calling this function.
//
// TODO_MAINNET(@Olshansk): Research if there's a way to limit who can call this function?
func (k Keeper) SettleSessionAccounting(
	ctx context.Context,
	claim *prooftypes.Claim,
) error {
	logger := k.Logger().With("method", "SettleSessionAccounting")

	settlementAmt := sdk.NewCoin("upokt", math.NewInt(0))
	isSuccessful := false
	// This is emitted only when the function returns.
	defer telemetry.EventSuccessCounter(
		"settle_session_accounting",
		func() float32 { return float32(settlementAmt.Amount.Int64()) },
		func() bool { return isSuccessful },
	)

	// Make sure the claim is not nil
	if claim == nil {
		logger.Error("received a nil claim")
		return types.ErrTokenomicsClaimNil
	}

	// Retrieve & validate the session header
	sessionHeader := claim.GetSessionHeader()
	if sessionHeader == nil {
		logger.Error("received a nil session header")
		return types.ErrTokenomicsSessionHeaderNil
	}
	if err := sessionHeader.ValidateBasic(); err != nil {
		logger.Error("received an invalid session header", "error", err)
		return types.ErrTokenomicsSessionHeaderInvalid
	}

	supplierAddr, err := sdk.AccAddressFromBech32(claim.GetSupplierAddress())
	if err != nil || supplierAddr == nil {
		return types.ErrTokenomicsSupplierAddressInvalid
	}

	applicationAddress, err := sdk.AccAddressFromBech32(sessionHeader.GetApplicationAddress())
	if err != nil || applicationAddress == nil {
		return types.ErrTokenomicsApplicationAddressInvalid
	}

	// Retrieve the sum of the root as a proxy into the amount of work done
	root := (smt.MerkleRoot)(claim.GetRootHash())

	// TODO_BLOCKER(@Olshansk): This check should be the responsibility of the SMST package
	// since it's used to get compute units from the root hash.
	if root == nil || len(root) != smt.SmstRootSizeBytes {
		logger.Error(fmt.Sprintf("received an invalid root hash of size: %d", len(root)))
		return types.ErrTokenomicsRootHashInvalid
	}
	claimComputeUnits := root.Sum()

	// Helpers for logging the same metadata throughout this function calls
	logger = logger.With(
		"compute_units", claimComputeUnits,
		"session_id", sessionHeader.GetSessionId(),
		"supplier", supplierAddr,
		"application", applicationAddress,
	)

	logger.Info("About to start session settlement accounting")

	// Retrieve the staked application record
	application, foundApplication := k.applicationKeeper.GetApplication(ctx, applicationAddress.String())
	if !foundApplication {
		logger.Warn(fmt.Sprintf("application for claim with address %q not found", applicationAddress))
		return types.ErrTokenomicsApplicationNotFound
	}

	logger.Info(fmt.Sprintf("About to start settling claim for %d compute units", claimComputeUnits))

	// Calculate the amount of tokens to mint & burn
	settlementAmt = k.getCoinFromComputeUnits(ctx, root)
	settlementAmtuPOKT := sdk.NewCoins(settlementAmt)

	logger.Info(fmt.Sprintf(
		"%d compute units equate to %s for session %s",
		claimComputeUnits,
		settlementAmt,
		sessionHeader.SessionId,
	))

	// NB: We are doing a mint & burn + transfer, instead of a simple transfer
	// of funds from the supplier to the application in order to enable second
	// order economic effects with more optionality. This could include funds
	// going to pnf, delegators, enabling bonuses/rebates, etc...

	// Mint new uPOKT to the supplier module account.
	// These funds will be transferred to the supplier below.
	if err := k.bankKeeper.MintCoins(
		ctx, suppliertypes.ModuleName, settlementAmtuPOKT,
	); err != nil {
		return types.ErrTokenomicsSupplierModuleMintFailed.Wrapf(
			"minting %s to the supplier module account: %v",
			settlementAmt,
			err,
		)
	}
	logger.Info(fmt.Sprintf("minted %s in the supplier module", settlementAmt))

	// Send the newley minted uPOKT from the supplier module account
	// to the supplier's account.
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, suppliertypes.ModuleName, supplierAddr, settlementAmtuPOKT,
	); err != nil {
		return types.ErrTokenomicsSupplierModuleMintFailed.Wrapf(
			"sending %s to supplier with address %s: %v",
			settlementAmt,
			supplierAddr,
			err,
		)
	}
	logger.Info(fmt.Sprintf("sent %s from the supplier module to the supplier account with address %q", settlementAmt, supplierAddr))

	// Verify that the application has enough uPOKT to pay for the services it consumed
	if application.Stake.IsLT(settlementAmt) {
		logger.Warn(fmt.Sprintf(
			"THIS SHOULD NEVER HAPPEN. Application with address %s needs to be charged more than it has staked: %v > %v",
			applicationAddress,
			settlementAmtuPOKT,
			application.Stake,
		))
		// TODO_MAINNET(@Olshansk, @RawthiL): The application was over-serviced in the last session so it basically
		// goes "into debt". Need to design a way to handle this when we implement
		// probabilistic proofs and add all the parameter logic. Do we touch the application balance?
		// Do we just let it go into debt? Do we penalize the application? Do we unstake it? Etc...
		settlementAmt = sdk.NewCoin("upokt", math.Int(application.Stake.Amount))
		settlementAmtuPOKT = sdk.NewCoins(settlementAmt)
		// TODO_BLOCKER(@Olshansk): The application should be immediately unstaked at this point in time
	}

	// Burn uPOKT from the application module account which was held in escrow
	// on behalf of the application account.
	if err := k.bankKeeper.BurnCoins(
		ctx, apptypes.ModuleName, settlementAmtuPOKT,
	); err != nil {
		return types.ErrTokenomicsApplicationModuleBurn.Wrapf("burning %s from the application module account: %v", settlementAmt, err)
	}
	logger.Info(fmt.Sprintf("burned %s from the application module account", settlementAmt))

	// Update the application's on-chain stake
	newAppStake, err := (*application.Stake).SafeSub(settlementAmt)
	if err != nil {
		return types.ErrTokenomicsApplicationNewStakeInvalid.Wrapf("application %q stake cannot be reduced to a negative amount %v", applicationAddress, newAppStake)
	}
	application.Stake = &newAppStake
	k.applicationKeeper.SetApplication(ctx, application)
	logger.Info(fmt.Sprintf("updated stake for application with address %q to %s", applicationAddress, newAppStake))

	isSuccessful = true
	return nil
}

func (k Keeper) getCoinFromComputeUnits(ctx context.Context, root smt.MerkleRoot) sdk.Coin {
	// Retrieve the existing tokenomics params
	params := k.GetParams(ctx)

	upokt := math.NewInt(int64(root.Sum() * params.ComputeUnitsToTokensMultiplier))
	return sdk.NewCoin("upokt", upokt)
}
