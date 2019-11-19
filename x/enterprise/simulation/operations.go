package simulation

import (
	"errors"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/unification-com/mainchain-cosmos/simapp/helpers"
	"github.com/unification-com/mainchain-cosmos/x/enterprise/internal/keeper"
	"github.com/unification-com/mainchain-cosmos/x/enterprise/internal/types"
)

// SimulateMsgRaisePurchaseOrder generates a MsgUndPurchaseOrder with random values
// nolint: funlen
func SimulateMsgRaisePurchaseOrder(ak auth.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {

		simAccount, _ := simulation.RandomAcc(r, accs)
		account := ak.GetAccount(ctx, simAccount.Address)

		coins := account.SpendableCoins(ctx.BlockTime())

		fees, err := simulation.RandomFees(r, ctx, coins)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		msg := types.NewMsgUndPurchaseOrder(
			simAccount.Address,
			sdk.NewInt64Coin(types.DefaultDenomination, int64(simulation.RandIntBetween(r, 100000000000, 1000000000000))),
		)

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			simAccount.PrivKey,
		)

		res := app.Deliver(tx)
		if !res.IsOK() {
			return simulation.NoOpMsg(types.ModuleName), nil, errors.New(res.Log)
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil

	}
}

func SimulateMsgProcessUndPurchaseOrder(ak auth.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {

		params := types.NewQueryPurchaseOrdersParams(1, 100, types.StatusRaised, sdk.AccAddress{})
		raisedPos := k.GetPurchaseOrdersFiltered(ctx, params)

		// needs to be sent specifically by the designated Ent account
		entAcc := GenerateEntSourceSimAccount()
		account := ak.GetAccount(ctx, entAcc.Address)

		rndPo := simulation.RandIntBetween(r, 0, len(raisedPos))
		po := raisedPos[rndPo]

		coins := account.SpendableCoins(ctx.BlockTime())

		fees, err := simulation.RandomFees(r, ctx, coins)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		msg := types.NewMsgProcessUndPurchaseOrder(po.PurchaseOrderID, randomDecision(r), entAcc.Address)

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			entAcc.PrivKey,
		)

		res := app.Deliver(tx)
		if !res.IsOK() {
			return simulation.NoOpMsg(types.ModuleName), nil, errors.New(res.Log)
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func randomDecision(r *rand.Rand) types.PurchaseOrderStatus {
	rnd := simulation.RandIntBetween(r, 0, 1)
	switch rnd {
	case 0:
		return types.StatusRejected
	case 1:
		return types.StatusAccepted
	default:
		return types.StatusAccepted
	}
}
