package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/mohammadreza-torkaman/wework/testutil/keeper"
	"github.com/mohammadreza-torkaman/wework/testutil/nullify"
	"github.com/mohammadreza-torkaman/wework/x/recruiter/keeper"
	"github.com/mohammadreza-torkaman/wework/x/recruiter/types"
	"github.com/stretchr/testify/require"
)

func createNJob(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Job {
	items := make([]types.Job, n)
	for i := range items {
		items[i].Id = keeper.AppendJob(ctx, items[i])
	}
	return items
}

func TestJobGet(t *testing.T) {
	keeper, ctx := keepertest.RecruiterKeeper(t)
	items := createNJob(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetJob(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestJobRemove(t *testing.T) {
	keeper, ctx := keepertest.RecruiterKeeper(t)
	items := createNJob(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveJob(ctx, item.Id)
		_, found := keeper.GetJob(ctx, item.Id)
		require.False(t, found)
	}
}

func TestJobGetAll(t *testing.T) {
	keeper, ctx := keepertest.RecruiterKeeper(t)
	items := createNJob(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllJob(ctx)),
	)
}

func TestJobCount(t *testing.T) {
	keeper, ctx := keepertest.RecruiterKeeper(t)
	items := createNJob(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetJobCount(ctx))
}
