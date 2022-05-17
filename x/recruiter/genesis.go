package recruiter

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mohammadreza-torkaman/wework/x/recruiter/keeper"
	"github.com/mohammadreza-torkaman/wework/x/recruiter/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the job
	for _, elem := range genState.JobList {
		k.SetJob(ctx, elem)
	}

	// Set job count
	k.SetJobCount(ctx, genState.JobCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.JobList = k.GetAllJob(ctx)
	genesis.JobCount = k.GetJobCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
