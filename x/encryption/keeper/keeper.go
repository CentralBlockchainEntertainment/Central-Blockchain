package keeper

import (
    "github.com/CentralBlockchainEntertainment/Central-Blockchain/x/encryption/types"
    "github.com/CentralBlockchainEntertainment/Central-Blockchain/x/encryption/aes_encryption"
    "github.com/CentralBlockchainEntertainment/Central-Blockchain/x/encryption/ip_encryption"
    
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/cosmos-sdk/store/prefix"
    "github.com/cosmos/cosmos-sdk/codec"
    paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Keeper defines the encryption module's keeper
type Keeper struct {
    storeKey     sdk.StoreKey
    cdc          codec.Codec
    paramstore   paramtypes.Subspace // Parameter store reference
}

// NewKeeper returns a new instance of the Keeper
func NewKeeper(storeKey sdk.StoreKey, cdc codec.Codec, ps paramtypes.Subspace) Keeper {
    // Initialize the parameter store and set default values if necessary
    if !ps.HasKeyTable() {
        ps = ps.WithKeyTable(types.ParamKeyTable())
    }

    return Keeper{
        storeKey:   storeKey,
        cdc:        cdc,
        paramstore: ps,
    }
}

// GetEncryptionStrength retrieves the encryption strength parameter from the store
func (k Keeper) GetEncryptionStrength(ctx sdk.Context) int {
    var strength int
    k.paramstore.Get(ctx, types.KeyEncryptionStrength, &strength)
    return strength
}

// SetParams sets the module parameters
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
    k.paramstore.SetParamSet(ctx, &params)
}

// GetParams retrieves the module parameters
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
    var params types.Params
    k.paramstore.GetParamSet(ctx, &params)
    return params
}

// EncryptData handles encryption of data using either AES or IP-based encryption,
// depending on the specified encryptionType. This function assumes aes_encryption.go and
// ip_encryption.go contain EncryptAES and EncryptWithIP methods.
func (k Keeper) EncryptData(ctx sdk.Context, data string, encryptionType string, ip string) (string, error) {
    var encryptedData string
    var err error

    switch encryptionType {
    case "AES":
        encryptedData, err = aes_encryption.EncryptAES(k.getAESKey(ctx), []byte(data))
    case "IP":
        encryptedData, err = ip_encryption.EncryptIP(data, k.getAESKey(ctx))
    default:
        return "", sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "unknown encryption type")
    }

    if err != nil {
        return "", sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "encryption failed")
    }

    // Store the encrypted data
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyAES())
    store.Set([]byte(data), []byte(encryptedData))

    return encryptedData, nil
}

// DecryptData handles decryption of encrypted data using either AES or IP-based decryption,
// depending on the specified encryptionType. Assumes DecryptAES and DecryptWithIP methods.
func (k Keeper) DecryptData(ctx sdk.Context, encryptedData string, encryptionType string, ip string) (string, error) {
    var decryptedData string
    var err error

    switch encryptionType {
    case "AES":
        decryptedData, err = aes_encryption.DecryptAES(k.getAESKey(ctx), []byte(encryptedData))
    case "IP":
        decryptedData, err = ip_encryption.DecryptIP(encryptedData, k.getAESKey(ctx))
    default:
        return "", sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "unknown encryption type")
    }

    if err != nil {
        return "", sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "decryption failed")
    }

    return decryptedData, nil
}

// SetModuleEnabled sets the module's enabled status in the store.
func (k Keeper) SetModuleEnabled(ctx sdk.Context, enabled bool) {
    store := ctx.KVStore(k.storeKey)
    var value byte
    if enabled {
        value = 1
    } else {
        value = 0
    }
    store.Set([]byte("moduleEnabled"), []byte{value})
}

// IsModuleEnabled retrieves the module's enabled status from the store.
func (k Keeper) IsModuleEnabled(ctx sdk.Context) bool {
    store := ctx.KVStore(k.storeKey)
    value := store.Get([]byte("moduleEnabled"))
    if value == nil {
        return false // Default to false if no value is found
    }
    return value[0] == 1
}

// getAESKey is a helper function that retrieves the AES encryption key.
// For simplicity, this example generates a static key.
// In a real implementation, key management would be more sophisticated.
func (k Keeper) getAESKey(ctx sdk.Context) []byte {
    // Here we return a hardcoded key for demonstration.
    // In production, consider fetching from a secure source or generating dynamically.
    return []byte("mysecureaeskey123") // Example key, replace with secure handling
}
