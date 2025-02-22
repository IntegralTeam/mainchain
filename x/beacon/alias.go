package beacon

import (
	"github.com/unification-com/mainchain/x/beacon/exported"
	"github.com/unification-com/mainchain/x/beacon/internal/ante"
	"github.com/unification-com/mainchain/x/beacon/internal/keeper"
	"github.com/unification-com/mainchain/x/beacon/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey

	DefaultParamspace = types.DefaultParamspace

	QueryParameters = keeper.QueryParameters
	QuerierRoute    = types.QuerierRoute
	QueryBeacons    = keeper.QueryBeacons

	RecordFee = types.RecordFee
	FeeDenom  = types.FeeDenom
)

var (
	NewKeeper       = keeper.NewKeeper
	NewQuerier      = keeper.NewQuerier
	NewGenesisState = types.NewGenesisState

	RegisterCodec = types.RegisterCodec
	ModuleCdc     = types.ModuleCdc

	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	NewQueryBeaconParams        = types.NewQueryBeaconParams
	NewMsgRegisterBeacon        = types.NewMsgRegisterBeacon
	NewMsgRecordBeaconTimestamp = types.NewMsgRecordBeaconTimestamp

	NewParamsRetriever = keeper.NewParamsRetriever

	// Errors
	ErrBeaconAlreadyRegistered        = types.ErrBeaconAlreadyRegistered
	ErrBeaconDoesNotExist             = types.ErrBeaconDoesNotExist
	ErrNotBeaconOwner                 = types.ErrNotBeaconOwner
	ErrBeaconTimestampAlreadyRecorded = types.ErrBeaconTimestampAlreadyRecorded
	ErrContentTooLarge                = types.ErrContentTooLarge
	ErrMissingData                    = types.ErrMissingData

	// Events
	EventTypeRegisterBeacon         = types.EventTypeRegisterBeacon
	EventTypeRecordBeaconTimestamp  = types.EventTypeRecordBeaconTimestamp
	AttributeValueCategory          = types.AttributeValueCategory
	AttributeKeyOwner               = types.AttributeKeyOwner
	AttributeKeyBeaconId            = types.AttributeKeyBeaconId
	AttributeKeyBeaconMoniker       = types.AttributeKeyBeaconMoniker
	AttributeKeyBeaconName          = types.AttributeKeyBeaconName
	AttributeKeyTimestampID         = types.AttributeKeyTimestampID
	AttributeKeyTimestampHash       = types.AttributeKeyTimestampHash
	AttributeKeyTimestampSubmitTime = types.AttributeKeyTimestampSubmitTime

	GetBeaconIDBytes    = types.GetBeaconIDBytes
	GetTimestampIDBytes = types.GetTimestampIDBytes

	CheckIsBeaconTx = exported.CheckIsBeaconTx

	DefaultParams = types.DefaultParams

	NewCorrectBeaconFeeDecorator = ante.NewCorrectBeaconFeeDecorator
)

type (
	Keeper = keeper.Keeper

	GenesisState = types.GenesisState

	Params                        = types.Params
	Beacon                        = types.Beacon
	BeaconTimestamp               = types.BeaconTimestamp
	BeaconExport                  = types.BeaconExport
	BeaconTimestampsGenesisExport = types.BeaconTimestampsGenesisExport

	QueryResBeacons               = types.QueryResBeacons
	QueryResBeaconTimestampHashes = types.QueryResBeaconTimestampHashes
	QueryBeaconParams             = types.QueryBeaconParams
	QueryBeaconTimestampParams    = types.QueryBeaconTimestampParams

	// Msgs
	MsgRegisterBeacon        = types.MsgRegisterBeacon
	MsgRecordBeaconTimestamp = types.MsgRecordBeaconTimestamp
)
