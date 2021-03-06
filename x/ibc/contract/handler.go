package contract

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/datachainlab/cross/x/ibc/contract/types"
	"github.com/datachainlab/cross/x/ibc/cross"
)

// NewHandler returns a handler
func NewHandler(k Keeper, contractHandler cross.ContractHandler) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case MsgContractCall:
			return handleContractCall(ctx, msg, k, contractHandler)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized message type: %T", msg)
		}
	}
}

func handleContractCall(ctx sdk.Context, msg MsgContractCall, k Keeper, contractHandler cross.ContractHandler) (*sdk.Result, error) {
	ctx = cross.WithSigners(ctx, msg.GetSigners())
	state, res, err := contractHandler.Handle(ctx, msg.Contract)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrFailedContractHandle, err.Error())
	}
	if err := state.CommitImmediately(); err != nil {
		return nil, sdkerrors.Wrap(types.ErrFailedCommitStore, err.Error())
	}
	res = contractHandler.OnCommit(ctx, res)
	data, err := k.MakeContractCallResponseData(res.GetData(), state.OPs())
	if err != nil { // internal error
		return nil, err
	}
	ctx.EventManager().EmitEvents(res.GetEvents())
	return &sdk.Result{Data: data, Events: ctx.EventManager().ABCIEvents()}, nil
}
