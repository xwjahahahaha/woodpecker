package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	// Attribute
	ErrAttributeNotExist = sdkerrors.Register(ModuleName, 1001, "Attribute not exist")
	ErrAttributeParameterCantEmpty = sdkerrors.Register(ModuleName, 1002, "ErrAttributeParameterCantEmpty")

	// MedicalHistory
	ErrMedicalHistoryParameterCantEmpty = sdkerrors.Register(ModuleName, 2002, "ErrMedicalHistoryParameterCantEmpty")

	// BodyIndex
	ErrBodyIndexParameterCantEmpty = sdkerrors.Register(ModuleName, 3002, "ErrBodyIndexParameterCantEmpty")
)
