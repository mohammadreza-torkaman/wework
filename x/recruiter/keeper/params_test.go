package keeper_test

import (
	"testing"

	testkeeper "github.com/mohammadreza-torkaman/wework/testutil/keeper"
	"github.com/mohammadreza-torkaman/wework/x/recruiter/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RecruiterKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
