package types

import (
    "fmt"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState represents the initial state of the encryption module at genesis.
type GenesisState struct {
    // Placeholder for configurable parameters, if any.
    Enabled bool `json:"enabled"` // Example parameter to enable/disable encryption module
}

// DefaultGenesis provides the default genesis state for the encryption module.
func DefaultGenesis() *GenesisState {
    return &GenesisState{
        Enabled: true, // Default to enabled
    }
}

// ValidateGenesis checks that the genesis state is valid.
func ValidateGenesis(data GenesisState) error {
    if !data.Enabled {
        return fmt.Errorf("encryption module must be enabled at genesis")
    }
    return nil
}

// InitGenesis initializes the encryption module’s state at genesis.
func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) {
    if data.Enabled {
        // Initialize module-specific data if required
    }
    k.SetModuleEnabled(ctx, data.Enabled) // Store Enabled status in the keeper, if required
}

// ExportGenesis exports the encryption module’s state to the genesis file.
func ExportGenesis(ctx sdk.Context, k Keeper) *GenesisState {
    enabled := k.IsModuleEnabled(ctx) // Fetch Enabled status from the keeper, if applicable
    return &GenesisState{
        Enabled: enabled,
    }
}
