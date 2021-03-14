package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Attribute struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
    Name string `json:"name" yaml:"name"`
    IdNumber string `json:"idNumber" yaml:"idNumber"`
    Address string `json:"address" yaml:"address"`
    Job string `json:"job" yaml:"job"`
    Nation string `json:"nation" yaml:"nation"`
}