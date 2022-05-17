package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		JobList: []Job{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in job
	jobIdMap := make(map[uint64]bool)
	jobCount := gs.GetJobCount()
	for _, elem := range gs.JobList {
		if _, ok := jobIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for job")
		}
		if elem.Id >= jobCount {
			return fmt.Errorf("job id should be lower or equal than the last id")
		}
		jobIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
