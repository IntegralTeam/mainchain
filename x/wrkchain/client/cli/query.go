package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/unification-com/mainchain-cosmos/x/wrkchain/internal/types"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	wrkchainQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the wrkchain module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	wrkchainQueryCmd.AddCommand(client.GetCommands(
		GetCmdWrkChain(storeKey, cdc),
		GetCmdWrkChainBlock(storeKey, cdc),
		GetCmdWrkChainBlockHashes(storeKey, cdc),
	)...)
	return wrkchainQueryCmd
}

// GetCmdWrkChain queries information about a wrkchain
func GetCmdWrkChain(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get [wrkchain id]",
		Short: "Query a WRKChain for given ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			wrkchainId := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/get/%s", queryRoute, wrkchainId), nil)
			if err != nil {
				fmt.Printf("could not find WRKChain - %s \n", wrkchainId)
				return nil
			}

			var out types.WrkChain
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdWrkChainBlock queries information about a wrkchain's recorded block hashes
func GetCmdWrkChainBlock(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-block [wrkchain id] [height]",
		Short: "Query a WRKChain for given ID and block height to retrieve recorded hashes for that block",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			wrkchainId := args[0]
			height := args[1]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/get-block/%s/%s", queryRoute, wrkchainId, height), nil)
			if err != nil {
				fmt.Printf("could not find WRKChain %s block hashes at height %s \n", wrkchainId, height)
				return nil
			}

			var out types.WrkChainBlock
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdWrkChainBlockHashes queries a list of all recorded block hashes for a WRKChain
func GetCmdWrkChainBlockHashes(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "blocks [wrkchain id]",
		Short: "Query a WRKChain for all hashes recorded to date",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			wrkchainId := args[0]
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/blocks/%s", queryRoute, wrkchainId), nil)
			if err != nil {
				fmt.Printf("could not get query block hashes\n")
				return nil
			}

			var out types.QueryResWrkChainBlockHashes
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
