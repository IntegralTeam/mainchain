package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"
	undtypes "github.com/unification-com/mainchain/types"
	"github.com/unification-com/mainchain/x/enterprise/internal/types"
	"strconv"
	"strings"
)

const (
	FlagNumLimit            = "limit"
	FlagPage                = "page"
	FlagPurchaseOrderStatus = "status"
	FlagPurchaser           = "purchaser"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	enterpriseTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Enterprise UND transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	enterpriseTxCmd.AddCommand(flags.PostCommands(
		GetCmdRaisePurchaseOrder(cdc),
		GetCmdProcessPurchaseOrder(cdc),
	)...)

	return enterpriseTxCmd
}

// GetCmdRegisterWrkChain is the CLI command for creating an Enterprise UND purchase order
func GetCmdRaisePurchaseOrder(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "purchase [amount]",
		Short: "Raise a new Enterprise UND purchase order",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Raise a new Enterprise UND purchase order
Example:
$ %s tx %s purchase 1000000000000%s --from wrktest
$ %s tx %s purchase 1%s --from wrktest
`,
				version.ClientName, types.ModuleName, undtypes.DefaultDenomination,
				version.ClientName, types.ModuleName, undtypes.BaseDenomination,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			amount, err := sdk.ParseCoin(args[0])
			if err != nil {
				return err
			}

			if amount.Denom == undtypes.BaseDenomination {
				// convert
				fromAmount := amount.Amount.String()
				converted, err := undtypes.ConvertUndDenomination(fromAmount, undtypes.BaseDenomination, undtypes.DefaultDenomination)
				if err != nil {
					return err
				}
				amount, err = sdk.ParseCoin(converted)
				if err != nil {
					return err
				}
			}

			if amount.Denom != undtypes.DefaultDenomination {
				return sdk.ErrInvalidCoins(fmt.Sprintf("denomination should be %s", undtypes.DefaultDenomination))
			}

			msg := types.NewMsgUndPurchaseOrder(cliCtx.GetFromAddress(), amount)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdProcessPurchaseOrder is the CLI command for processing an Enterprise UND purchase order
func GetCmdProcessPurchaseOrder(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "process [purchase_order_id] [decision]",
		Short: "Process an Enterprise UND purchase order",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Process an Enterprise UND purchase order
Example:
$ %s tx %s process 24 accept --from ent
$ %s tx %s process 24 reject --from ent
`,
				version.ClientName, types.ModuleName, version.ClientName, types.ModuleName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			purchaseOrderId, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			decision, err := types.PurchaseOrderStatusFromString(args[1])
			if err != nil {
				return err
			}

			if !types.ValidPurchaseOrderAcceptRejectStatus(decision) {
				return types.ErrInvalidDecision(types.DefaultCodespace, "decision should be accept or reject")
			}

			msg := types.NewMsgProcessUndPurchaseOrder(uint64(purchaseOrderId), decision, cliCtx.GetFromAddress())
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
