package types

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
    TypeMsgEncrypt = "encrypt"
    TypeMsgDecrypt = "decrypt"
)

// MsgEncrypt represents a request to encrypt data
type MsgEncrypt struct {
    Creator sdk.AccAddress `json:"creator"`
    Data    string         `json:"data"`
}

// Route returns the module's message route for MsgEncrypt
func (msg MsgEncrypt) Route() string { return RouterKey }

// Type returns the type of the message
func (msg MsgEncrypt) Type() string { return TypeMsgEncrypt }

// GetSigners returns the address of the message signer
func (msg MsgEncrypt) GetSigners() []sdk.AccAddress {
    return []sdk.AccAddress{msg.Creator}
}

// ValidateBasic performs basic validation on the MsgEncrypt message
func (msg MsgEncrypt) ValidateBasic() error {
    if msg.Creator.Empty() {
        return errors.Wrap(errors.ErrInvalidAddress, "creator can't be empty")
    }
    if len(msg.Data) == 0 {
        return errors.Wrap(errors.ErrUnknownRequest, "data to encrypt can't be empty")
    }
    return nil
}

// MsgDecrypt represents a request to decrypt data
type MsgDecrypt struct {
    Creator sdk.AccAddress `json:"creator"`
    Data    string         `json:"data"`
}

// Route returns the module's message route for MsgDecrypt
func (msg MsgDecrypt) Route() string { return RouterKey }

// Type returns the type of the message
func (msg MsgDecrypt) Type() string { return TypeMsgDecrypt }

// GetSigners returns the address of the message signer
func (msg MsgDecrypt) GetSigners() []sdk.AccAddress {
    return []sdk.AccAddress{msg.Creator}
}

// ValidateBasic performs basic validation on the MsgDecrypt message
func (msg MsgDecrypt) ValidateBasic() error {
    if msg.Creator.Empty() {
        return errors.Wrap(errors.ErrInvalidAddress, "creator can't be empty")
    }
    if len(msg.Data) == 0 {
        return errors.Wrap(errors.ErrUnknownRequest, "data to decrypt can't be empty")
    }
    return nil
}
