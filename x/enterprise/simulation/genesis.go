package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/unification-com/mainchain-cosmos/x/enterprise/internal/types"
)

// Simulation parameter constants
const (
	EnterpriseSourceAddress = "ent_source"
)

// RandomizedGenState generates a random GenesisState for enterprise
func RandomizedGenState(simState *module.SimulationState) {
	startingPurchaseOrderID := uint64(simState.Rand.Intn(100))

    var entAddress sdk.AccAddress
	simState.AppParams.GetOrGenerate(
		simState.Cdc, EnterpriseSourceAddress, &entAddress, simState.Rand,
		func(r *rand.Rand) { entAddress = GetEntSourceAddress() },
	)

	entGenesis := types.NewGenesisState(
		types.NewParams(entAddress),
		startingPurchaseOrderID,
		sdk.NewInt64Coin(types.DefaultDenomination, 0),
	)

	fmt.Printf("Selected randomly generated enterprise parameters:\n%s\n", codec.MustMarshalJSONIndent(simState.Cdc, entGenesis))
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(entGenesis)
}
