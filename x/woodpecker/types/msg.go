package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Transactions messages must fulfill the Msg
// 实现接口必须实现所有的函数
type Msg interface {
	// Return the message type.
	// Must be alphanumeric or empty.
	// 返回消息类型， 必须是字母或者空
	Type() string

	// Returns a human-readable string for the message, intended for utilization
	// within tags
	// 返回消息的可读string,用于在标签中使用
	Route() string

	// ValidateBasic does a simple validation check that
	// doesn't require access to any other information.
	// 做一些基本的验证, 不需要访问其他信息
	ValidateBasic() error

	// Get the canonical byte representation of the Msg.
	// 获取Msg的规范字节表示即字节数组
	GetSignBytes() []byte

	// Signers returns the addrs of signers that must sign.
	// CONTRACT: All signatures must be present to be valid.
	// CONTRACT: Returns addrs in some deterministic order.
	// 返回必须签名的签名者地址集合， 所有的签名必须在当下还有效， 返回的签名集合会以某种确定的顺序
	GetSigners() []sdk.AccAddress
}