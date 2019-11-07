package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/unification-com/mainchain-cosmos/x/enterprise/internal/types"
	"strconv"
)

const (
	QueryParameters       = "params"
	QueryPurchaseOrders   = "get-all-pos"
	QueryGetPurchaseOrder = "get"
	QueryGetLocked        = "locked"
	QueryTotalLocked      = "total-locked"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryParameters:
			return queryParams(ctx, keeper)
		case QueryPurchaseOrders:
			return queryPurchaseOrders(ctx, keeper)
		case QueryGetPurchaseOrder:
			return queryPurchaseOrderById(ctx, path[1:], keeper)
		case QueryGetLocked:
			return queryLockedUndByAddress(ctx, path[1:], keeper)
		case QueryTotalLocked:
			return queryTotalLocked(ctx, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown wrkchain query endpoint")
		}
	}
}

func queryParams(ctx sdk.Context, k Keeper) ([]byte, sdk.Error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(k.cdc, params)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to marshal JSON", err.Error()))
	}

	return res, nil
}

func queryPurchaseOrders(ctx sdk.Context, k Keeper) ([]byte, sdk.Error) {
	posIterator := k.GetAllPurchaseOrdersIterator(ctx)
	var pos []types.EnterpriseUndPurchaseOrder

	for ; posIterator.Valid(); posIterator.Next() {
		var po types.EnterpriseUndPurchaseOrder
		k.cdc.MustUnmarshalBinaryBare(posIterator.Value(), &po)
		pos = append(pos, po)
	}

	res, err := codec.MarshalJSONIndent(k.cdc, pos)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

func queryPurchaseOrderById(ctx sdk.Context, path []string, k Keeper) ([]byte, sdk.Error) {
	purchaseOrderId, err := strconv.Atoi(path[0])

	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to parse PurchaseOrderID", err.Error()))
	}

	purchaseOrder := k.GetPurchaseOrder(ctx, uint64(purchaseOrderId))

	res, err := codec.MarshalJSONIndent(k.cdc, purchaseOrder)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

func queryLockedUndByAddress(ctx sdk.Context, path []string, k Keeper) ([]byte, sdk.Error) {
	address, err := sdk.AccAddressFromBech32(path[0])
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to parse address", err.Error()))
	}

	lockedUnd := k.GetLockedUndForAccount(ctx, address)

	res, err := codec.MarshalJSONIndent(k.cdc, lockedUnd)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

func queryTotalLocked(ctx sdk.Context, k Keeper) ([]byte, sdk.Error) {

	totalLocked := k.GetTotalLockedUnd(ctx)

	res, err := codec.MarshalJSONIndent(k.cdc, totalLocked)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
