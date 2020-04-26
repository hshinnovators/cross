package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/datachainlab/cross/x/ibc/cross"
)

var _ sdk.Msg = (*MsgContractCall)(nil)

type MsgContractCall struct {
	Sender         sdk.AccAddress
	Signers        []sdk.AccAddress       `json:"signers" yaml:"signers"`
	CallInfo       cross.ContractCallInfo `json:"call_info" yaml:"call_info"`
	StateCondition cross.StateCondition   `json:"state_condition" yaml:"state_condition"`
}

func NewMsgContractCall(sender sdk.AccAddress, signers []sdk.AccAddress, callInfo cross.ContractCallInfo, cond cross.StateCondition) MsgContractCall {
	return MsgContractCall{
		Sender:         sender,
		Signers:        signers,
		CallInfo:       callInfo,
		StateCondition: cond,
	}
}

func (MsgContractCall) Route() string {
	return RouterKey
}

func (MsgContractCall) Type() string {
	return "contract_contractcall"
}

func (msg MsgContractCall) ValidateBasic() error {
	return nil
}

func (msg MsgContractCall) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements sdk.Msg
// GetSigners returns the addresses that must sign the transaction.
// Addresses are returned in a deterministic order.
// Duplicate addresses will be omitted.
func (msg MsgContractCall) GetSigners() []sdk.AccAddress {
	seen := map[string]bool{}
	signers := []sdk.AccAddress{msg.Sender}
	for _, addr := range msg.Signers {
		if !seen[addr.String()] {
			signers = append(signers, addr)
			seen[addr.String()] = true
		}
	}
	return signers
}
