package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/mohammadreza-torkaman/wework/testutil/keeper"
	"github.com/mohammadreza-torkaman/wework/x/recruiter/keeper"
	"github.com/mohammadreza-torkaman/wework/x/recruiter/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.RecruiterKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
