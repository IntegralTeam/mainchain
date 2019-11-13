package enterprise

import (
	"github.com/unification-com/mainchain-cosmos/x/enterprise/internal/keeper"
	"github.com/unification-com/mainchain-cosmos/x/enterprise/internal/types"
)

const (
	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace

	DefaultCodespace = types.DefaultCodespace

	DefaultDenomination = types.DefaultDenomination

	QuerierRoute          = types.QuerierRoute
	QueryParameters       = keeper.QueryParameters
	QueryPurchaseOrders   = keeper.QueryPurchaseOrders
	QueryGetPurchaseOrder = keeper.QueryGetPurchaseOrder

	StatusNil       = types.StatusNil
	StatusRaised    = types.StatusRaised
	StatusAccepted  = types.StatusAccepted
	StatusRejected  = types.StatusRejected
	StatusCompleted = types.StatusCompleted
)

var (
	NewKeeper           = keeper.NewKeeper
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	ModuleCdc           = types.ModuleCdc

	// Events
	EventTypeRaisePurchaseOrder   = types.EventTypeRaisePurchaseOrder
	EventTypeProcessPurchaseOrder = types.EventTypeProcessPurchaseOrder
	EventTypeUndPurchaseComplete  = types.EventTypeUndPurchaseComplete
	AttributeValueCategory        = types.AttributeValueCategory
	AttributeKeyPurchaseOrderID   = types.AttributeKeyPurchaseOrderID
	AttributeKeyPurchaser         = types.AttributeKeyPurchaser
	AttributeKeyAmount            = types.AttributeKeyAmount
	AttributeKeyDecision          = types.AttributeKeyDecision

	// Key functions
	GetPurchaseOrderIDBytes = types.GetPurchaseOrderIDBytes

	ValidPurchaseOrderAcceptRejectStatus = types.ValidPurchaseOrderAcceptRejectStatus

	NewLockedUndRetriever   = keeper.NewLockedUndRetriever
	NewTotalSupplyRetriever = keeper.NewTotalSupplyRetriever

	RegisterInvariants            = keeper.RegisterInvariants
	AllInvariants                 = keeper.AllInvariants

	NewQueryPurchaseOrdersParams = types.NewQueryPurchaseOrdersParams

	// Errors
	ErrInvalidDecision = types.ErrInvalidDecision
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
	Params       = types.Params

	QueryPurchaseOrdersParams = types.QueryPurchaseOrdersParams

	// Msgs
	MsgPurchaseUnd             = types.MsgUndPurchaseOrder
	MsgProcessUndPurchaseOrder = types.MsgProcessUndPurchaseOrder

	// Structs
	EnterpriseUndPurchaseOrder = types.EnterpriseUndPurchaseOrder
	UndSupply                  = types.UndSupply
	PurchaseOrderStatus        = types.PurchaseOrderStatus
)
