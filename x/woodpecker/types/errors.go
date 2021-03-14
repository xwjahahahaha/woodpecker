package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	// Attribute
	ErrAttributeNotExist = sdkerrors.Register(ModuleName, 1001, "Attribute not exist")
	ErrAttributeParameterCantEmpty = sdkerrors.Register(ModuleName, 1002, "Attribute's parameter Can't Empty")
)
