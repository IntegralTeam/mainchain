package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/unification-com/mainchain/x/enterprise/internal/types"
)

// Simulation parameter constants
const (
	EnterpriseSourceAddress = "ent_signers"
)

// RandomizedGenState generates a random GenesisState for enterprise
func RandomizedGenState(simState *module.SimulationState) {
	startingPurchaseOrderID := uint64(simState.Rand.Intn(100))

	var entAddress sdk.AccAddress
	simState.AppParams.GetOrGenerate(
		simState.Cdc, EnterpriseSourceAddress, &entAddress, simState.Rand,
		func(r *rand.Rand) { entAddress = GetEntSourceAddress() },
	)

	var entAddresses []sdk.AccAddress
	entAddresses = append(entAddresses, entAddress)

	// NOTE: for simulation, we're using sdk.DefaultBondDenom ("stake"), since "stake" is hard-coded
	// into the SDK's module simulation functions
	entGenesis := types.NewGenesisState(
		types.NewParams(entAddresses, sdk.DefaultBondDenom, 1, 3600),
		startingPurchaseOrderID,
		sdk.NewInt64Coin(sdk.DefaultBondDenom, 0),
	)

	fmt.Printf("Selected randomly generated enterprise parameters:\n%s\n", codec.MustMarshalJSONIndent(simState.Cdc, entGenesis))
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(entGenesis)
}
