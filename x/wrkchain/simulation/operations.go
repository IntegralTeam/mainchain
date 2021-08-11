package simulation

//import (
//	"math/rand"
//
//	"github.com/cosmos/cosmos-sdk/baseapp"
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	"github.com/cosmos/cosmos-sdk/x/auth"
//	"github.com/cosmos/cosmos-sdk/x/simulation"
//	"github.com/unification-com/mainchain/simapp/helpers"
//	"github.com/unification-com/mainchain/x/wrkchain/internal/keeper"
//	"github.com/unification-com/mainchain/x/wrkchain/internal/types"
//)
//
//// SimulateMsgRegisterWrkChain generates a MsgRegisterWrkChain with random values
//// nolint: funlen
//func SimulateMsgRegisterWrkChain(ak auth.AccountKeeper, k keeper.Keeper) simulation.Operation {
//	return func(
//		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
//	) (simulation.OperationMsg, []simulation.FutureOperation, error) {
//
//		simAccount, _ := simulation.RandomAcc(r, accs)
//		account := ak.GetAccount(ctx, simAccount.Address)
//
//		moniker := simulation.RandStringOfLength(r, 64)
//		name := simulation.RandStringOfLength(r, 128)
//
//		fees := k.GetRegistrationFeeAsCoins(ctx)
//
//		coins := account.SpendableCoins(ctx.BlockTime())
//
//		_, hasNeg := coins.SafeSub(fees)
//
//		if hasNeg {
//			return simulation.NoOpMsg(types.ModuleName), nil, nil
//		}
//
//		msg := types.NewMsgRegisterWrkChain(
//			moniker,
//			simulation.RandStringOfLength(r, 66),
//			name,
//			"geth",
//			simAccount.Address,
//		)
//
//		tx := helpers.GenTx(
//			[]sdk.Msg{msg},
//			fees,
//			chainID,
//			[]uint64{account.GetAccountNumber()},
//			[]uint64{account.GetSequence()},
//			simAccount.PrivKey,
//		)
//
//		_, _, err := app.Deliver(tx)
//		if err != nil {
//			return simulation.NoOpMsg(types.ModuleName), nil, err
//		}
//
//		return simulation.NewOperationMsg(msg, true, ""), nil, nil
//
//	}
//}
//
//// SimulateMsgRecordWrkChainBlock generates a MsgRecordWrkChainBlock with random values
//// nolint: funlen
//func SimulateMsgRecordWrkChainBlock(ak auth.AccountKeeper, k keeper.Keeper) simulation.Operation {
//	return func(
//		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
//	) (simulation.OperationMsg, []simulation.FutureOperation, error) {
//
//		// randomly select a WRKChain
//		wrkChains := k.GetAllWrkChains(ctx)
//		if len(wrkChains) == 0 {
//			// nothing registered yet
//			return simulation.NoOpMsg(types.ModuleName), nil, nil
//		}
//
//		rWc := 0
//		if len(wrkChains) > 1 {
//			rWc = rand.Intn(len(wrkChains) - 1)
//		}
//		wrkChain := wrkChains[rWc]
//
//		ownerAddr := wrkChain.Owner
//		var simAccount simulation.Account
//
//		for _, ac := range accs {
//			if ac.Address.String() == ownerAddr.String() {
//				simAccount = ac
//			}
//		}
//
//		account := ak.GetAccount(ctx, ownerAddr)
//
//		height := wrkChain.LastBlock + 1
//
//		msg := types.NewMsgRecordWrkChainBlock(
//			wrkChain.WrkChainID,
//			height,
//			simulation.RandStringOfLength(r, 66),
//			simulation.RandStringOfLength(r, 66),
//			simulation.RandStringOfLength(r, 66),
//			simulation.RandStringOfLength(r, 66),
//			simulation.RandStringOfLength(r, 66),
//			ownerAddr,
//		)
//
//		fees := k.GetRecordFeeAsCoins(ctx)
//
//		coins := account.SpendableCoins(ctx.BlockTime())
//		_, hasNeg := coins.SafeSub(fees)
//
//		if hasNeg {
//			return simulation.NoOpMsg(types.ModuleName), nil, nil
//		}
//
//		tx := helpers.GenTx(
//			[]sdk.Msg{msg},
//			fees,
//			chainID,
//			[]uint64{account.GetAccountNumber()},
//			[]uint64{account.GetSequence()},
//			simAccount.PrivKey,
//		)
//
//		_, _, err := app.Deliver(tx)
//		if err != nil {
//			return simulation.NoOpMsg(types.ModuleName), nil, err
//		}
//
//		return simulation.NewOperationMsg(msg, true, ""), nil, nil
//
//	}
//}
