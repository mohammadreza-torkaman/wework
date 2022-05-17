package recruiter_test

import (
	"testing"

	keepertest "github.com/mohammadreza-torkaman/wework/testutil/keeper"
	"github.com/mohammadreza-torkaman/wework/testutil/nullify"
	"github.com/mohammadreza-torkaman/wework/x/recruiter"
	"github.com/mohammadreza-torkaman/wework/x/recruiter/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		JobList: []types.Job{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		JobCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RecruiterKeeper(t)
	recruiter.InitGenesis(ctx, *k, genesisState)
	got := recruiter.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.JobList, got.JobList)
	require.Equal(t, genesisState.JobCount, got.JobCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
