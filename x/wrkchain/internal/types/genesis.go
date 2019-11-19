package types

import (
	"fmt"
)

// GenesisState - enterprise state
type GenesisState struct {
	Params             Params           `json:"params" yaml:"params"`                             // wrkchain params
	StartingWrkChainID uint64           `json:"starting_wrkchain_id" yaml:"starting_wrkchain_id"` // should be 1
	WrkChains          []WrkChainExport `json:"registered_wrkchains"`
}

type WrkChainExport struct {
	WrkChain       WrkChain
	WrkChainBlocks []WrkChainBlock
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState(params Params, startingWrkChainID uint64) GenesisState {
	return GenesisState{
		Params:             params,
		StartingWrkChainID: startingWrkChainID,
		WrkChains:          nil,
	}
}

// DefaultGenesisState creates a default GenesisState object
func DefaultGenesisState() GenesisState {
	return GenesisState{
		Params:             DefaultParams(),
		StartingWrkChainID: DefaultStartingWrkChainID,
	}
}

// ValidateGenesis validates the provided genesis state to ensure the
// expected invariants holds.
func ValidateGenesis(data GenesisState) error {
	err := ValidateParams(data.Params)
	if err != nil {
		return err
	}

	for _, record := range data.WrkChains {
		if record.WrkChain.Owner == nil {
			return fmt.Errorf("Invalid WrkChain: Owner: %s. Error: Missing Owner", record.WrkChain.Owner)
		}
		if record.WrkChain.WrkChainID == 0 {
			return fmt.Errorf("Invalid WrkChain: Moniker: %d. Error: Missing ID", record.WrkChain.WrkChainID)
		}
		if record.WrkChain.GenesisHash == "" {
			return fmt.Errorf("Invalid WrkChain: GenesisHash: %s. Error: Missing Genesis Timestamp", record.WrkChain.GenesisHash)
		}
	}
	return nil
}
