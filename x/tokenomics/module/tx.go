package tokenomics

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/pokt-network/poktroll/x/tokenomics/types"
)

// GetTxCmd returns the transaction commands for this module
// TODO_TECHDEBT(#370, @bryanchriswhite): remove if custom query commands are consolidated into AutoCLI.
func (am AppModule) GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdUpdateParams())
	// this line is used by starport scaffolding # 1

	return cmd
}
