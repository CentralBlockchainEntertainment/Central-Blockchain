package types_test

import (
    "testing"

    "github.com/CentralBlockchainEntertainment/Central-Blockchain/x/encryption/types"
    "github.com/stretchr/testify/require"
)

func TestDefaultGenesis(t *testing.T) {
    defaultGenesis := types.DefaultGenesis()
    require.True(t, defaultGenesis.Enabled, "default genesis should have Enabled set to true")
}

func TestValidateGenesis(t *testing.T) {
    // Valid genesis state
    validGenesis := types.GenesisState{Enabled: true}
    err := types.ValidateGenesis(validGenesis)
    require.NoError(t, err, "valid genesis state should pass validation")

    // Invalid genesis state
    invalidGenesis := types.GenesisState{Enabled: false}
    err = types.ValidateGenesis(invalidGenesis)
    require.Error(t, err, "invalid genesis state should fail validation")
}
