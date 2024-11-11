package types

import (
    "fmt"
    sdk "github.com/cosmos/cosmos-sdk/types"
    paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter keys
var (
    KeyEncryptionStrength = []byte("EncryptionStrength")
)

// Default values for parameters
const (
    DefaultEncryptionStrength = 256 // Default strength, e.g., 256-bit encryption
)

// Params defines the parameters for the encryption module
type Params struct {
    EncryptionStrength int `json:"encryption_strength" yaml:"encryption_strength"`
}

// NewParams creates a new Params object
func NewParams(encryptionStrength int) Params {
    return Params{
        EncryptionStrength: encryptionStrength,
    }
}

// DefaultParams returns the default parameters for the encryption module
func DefaultParams() Params {
    return NewParams(DefaultEncryptionStrength)
}

// Validate checks that the parameters have valid values
func (p Params) Validate() error {
    if p.EncryptionStrength != 128 && p.EncryptionStrength != 256 {
        return fmt.Errorf("invalid encryption strength: %d", p.EncryptionStrength)
    }
    return nil
}

// ParamKeyTable returns the parameter key table for the encryption module
func ParamKeyTable() paramtypes.KeyTable {
    return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs returns the parameter set pairs for the encryption module
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
    return paramtypes.ParamSetPairs{
        paramtypes.NewParamSetPair(KeyEncryptionStrength, &p.EncryptionStrength, validateEncryptionStrength),
    }
}

// Validation function for EncryptionStrength
func validateEncryptionStrength(i interface{}) error {
    v, ok := i.(int)
    if !ok {
        return fmt.Errorf("invalid parameter type: %T", i)
    }
    if v != 128 && v != 256 {
        return fmt.Errorf("invalid encryption strength: %d", v)
    }
    return nil
}
