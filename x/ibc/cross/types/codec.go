package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgInitiate{}, "cross/MsgInitiate", nil)
	cdc.RegisterConcrete(CrossChainTransaction{}, "cross/CrossChainTransaction", nil)
	cdc.RegisterConcrete(CrossChainTransactions{}, "cross/CrossChainTransactions", nil)
	cdc.RegisterConcrete(ChannelInfo{}, "cross/ChannelInfo", nil)
	cdc.RegisterInterface((*PacketData)(nil), nil)
	cdc.RegisterConcrete(PacketDataPrepare{}, "cross/PacketDataPrepare", nil)
	cdc.RegisterConcrete(PacketDataPrepareResult{}, "cross/PacketDataPrepareResult", nil)
	cdc.RegisterConcrete(PacketDataCommit{}, "cross/PacketDataCommit", nil)
	cdc.RegisterConcrete(PacketDataAckCommit{}, "cross/PacketDataAckCommit", nil)
	cdc.RegisterInterface((*OP)(nil), nil)
	cdc.RegisterInterface((*ContractHandlerResult)(nil), nil)
}
