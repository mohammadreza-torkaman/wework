package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mohammadreza-torkaman/wework/x/recruiter/types"
)

func (k msgServer) SubmitJob(goCtx context.Context, msg *types.MsgSubmitJob) (*types.MsgSubmitJobResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitJobResponse{}, nil
}
