package encryption_test

import (
    "testing"

    "github.com/CentralBlockchainEntertainment/Central-Blockchain/x/encryption"
    "github.com/CentralBlockchainEntertainment/Central-Blockchain/x/encryption/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/stretchr/testify/require"
)

func TestAppModule(t *testing.T) {
    appModule := encryption.AppModule{}

    // Check the module name
    require.Equal(t, types.ModuleName, appModule.Name(), "module name should match")

    // Check that the module can be routed correctly
    route := appModule.Route()
    require.Equal(t, types.RouterKey, route.Path(), "route path should match module's RouterKey")
}
