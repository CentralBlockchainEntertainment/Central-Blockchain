package types_test

import (
    "testing"

    "github.com/CentralBlockchainEntertainment/Central-Blockchain/x/encryption/types"
    "github.com/stretchr/testify/require"
)

func TestCodec(t *testing.T) {
    cdc := types.ModuleCdc

    // Test MsgEncrypt
    originalEncryptMsg := types.MsgEncrypt{
        Creator: []byte("test_creator"),
        Data:    "test data",
    }
    bz, err := cdc.MarshalJSON(&originalEncryptMsg)
    require.NoError(t, err, "encoding MsgEncrypt should not produce an error")

    var decodedEncryptMsg types.MsgEncrypt
    err = cdc.UnmarshalJSON(bz, &decodedEncryptMsg)
    require.NoError(t, err, "decoding MsgEncrypt should not produce an error")
    require.Equal(t, originalEncryptMsg, decodedEncryptMsg, "encoded and decoded MsgEncrypt should be equal")

    // Test MsgDecrypt
    originalDecryptMsg := types.MsgDecrypt{
        Creator: []byte("test_creator"),
        Data:    "encrypted data",
    }
    bz, err = cdc.MarshalJSON(&originalDecryptMsg)
    require.NoError(t, err, "encoding MsgDecrypt should not produce an error")

    var decodedDecryptMsg types.MsgDecrypt
    err = cdc.UnmarshalJSON(bz, &decodedDecryptMsg)
    require.NoError(t, err, "decoding MsgDecrypt should not produce an error")
    require.Equal(t, originalDecryptMsg, decodedDecryptMsg, "encoded and decoded MsgDecrypt should be equal")
}
