package types_test

import (
    "testing"

    "github.com/CentralBlockchainEntertainment/Central-Blockchain/x/encryption/types"
    "github.com/stretchr/testify/require"
)

func TestParamsValidation(t *testing.T) {
    // Valid parameter set
    validParams := types.Params{
        EncryptionStrength: 256,
    }
    require.NoError(t, validParams.Validate(), "valid parameters should pass validation")

    // Invalid encryption strength
    invalidParams := types.Params{
        EncryptionStrength: 512, // Unsupported encryption strength
    }
    require.Error(t, invalidParams.Validate(), "invalid parameters should fail validation")
}
