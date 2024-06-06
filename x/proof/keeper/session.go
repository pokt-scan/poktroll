package keeper

import (
	"context"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/pokt-network/poktroll/x/proof/types"
	sessiontypes "github.com/pokt-network/poktroll/x/session/types"
	"github.com/pokt-network/poktroll/x/shared"
	sharedtypes "github.com/pokt-network/poktroll/x/shared/types"
)

type msgWithSessionAndSupplier interface {
	GetSessionHeader() *sessiontypes.SessionHeader
	GetSupplierAddress() string
}

// queryAndValidateSessionHeader ensures that a session with the sessionID of the given session
// header exists and that this session includes the supplier with the given address.
// It returns a session which is hydrated with the on-chain session data.
func (k msgServer) queryAndValidateSessionHeader(
	ctx context.Context,
	msg msgWithSessionAndSupplier,
) (*sessiontypes.Session, error) {
	logger := k.Logger().With("method", "queryAndValidateSessionHeader")

	sessionHeader := msg.GetSessionHeader()
	supplierAddr := msg.GetSupplierAddress()

	sessionReq := &sessiontypes.QueryGetSessionRequest{
		ApplicationAddress: sessionHeader.GetApplicationAddress(),
		Service:            sessionHeader.GetService(),
		BlockHeight:        sessionHeader.GetSessionStartBlockHeight(),
	}

	// Get the on-chain session for the ground-truth against which the given
	// session header is to be validated.
	sessionRes, err := k.Keeper.sessionKeeper.GetSession(ctx, sessionReq)
	if err != nil {
		return nil, err
	}
	onChainSession := sessionRes.GetSession()

	logger.
		With(
			"session_id", onChainSession.GetSessionId(),
			"session_end_height", sessionHeader.GetSessionEndBlockHeight(),
			"supplier", supplierAddr,
		).
		Debug("got sessionId for proof")

	// Ensure that the given session header's session ID matches the on-chain onChainSession ID.
	if sessionHeader.GetSessionId() != onChainSession.GetSessionId() {
		return nil, types.ErrProofInvalidSessionId.Wrapf(
			"session ID does not match on-chain session ID; expected %q, got %q",
			onChainSession.GetSessionId(),
			sessionHeader.GetSessionId(),
		)
	}

	// NB: it is redundant to assert that the service ID in the request matches the
	// on-chain session service ID because the session is queried using the service
	// ID as a parameter. Either a different session (i.e. different session ID)
	// or an error would be returned depending on whether an application/supplier
	// pair exists for the given service ID or not, respectively.

	// Ensure the given supplier is in the onChainSession supplier list.
	if isSupplerFound := foundSupplier(
		sessionRes.GetSession().GetSuppliers(),
		supplierAddr,
	); !isSupplerFound {
		return nil, types.ErrProofNotFound.Wrapf(
			"supplier address %q not found in session ID %q",
			supplierAddr,
			sessionHeader.GetSessionId(),
		)
	}

	return onChainSession, nil
}

// validateClaimWindow returns an error if the given session is not eligible for claiming.
// It *assumes* that the msg's session header is a valid on-chain session with correct
// height fields. First call #queryAndValidateSessionHeader to ensure any user-provided
// session header is valid and correctly hydrated.
func (k msgServer) validateClaimWindow(
	ctx context.Context,
	msg *types.MsgCreateClaim,
) error {
	logger := k.Logger().With("method", "validateClaimWindow")
	sessionHeader := msg.GetSessionHeader()
	sharedParams := k.sharedKeeper.GetParams(ctx)

	sessionEndHeight := sessionHeader.GetSessionEndBlockHeight()

	// Get the claim window open and close heights for the given session header.
	claimWindowOpenHeight := shared.GetClaimWindowOpenHeight(&sharedParams, sessionEndHeight)
	claimWindowCloseHeight := shared.GetClaimWindowCloseHeight(&sharedParams, sessionEndHeight)

	// Get the current block height.
	sdkCtx := cosmostypes.UnwrapSDKContext(ctx)
	currentHeight := sdkCtx.BlockHeight()

	// Ensure the current block height is AFTER the claim window open height.
	if currentHeight < claimWindowOpenHeight {
		return types.ErrProofClaimOutsideOfWindow.Wrapf(
			"current block height %d is less than session claim window open height %d",
			currentHeight,
			claimWindowOpenHeight,
		)
	}

	// Ensure the current block height is BEFORE the claim window close height.
	if currentHeight > claimWindowCloseHeight {
		return types.ErrProofClaimOutsideOfWindow.Wrapf(
			"current block height %d is greater than session claim window close height %d",
			currentHeight,
			claimWindowCloseHeight,
		)
	}

	logger.
		With(
			"current_height", currentHeight,
			"session_end_height", sessionEndHeight,
			"claim_window_open_height", claimWindowOpenHeight,
			"claim_window_close_height", claimWindowCloseHeight,
			"supplier_addr", msg.GetSupplierAddress(),
		).
		Debug("validated claim window")

	return nil
}

// foundSupplier ensures that the given supplier address is in the given list of suppliers.
func foundSupplier(suppliers []*sharedtypes.Supplier, supplierAddr string) bool {
	for _, supplier := range suppliers {
		if supplier.Address == supplierAddr {
			return true
		}
	}
	return false
}
