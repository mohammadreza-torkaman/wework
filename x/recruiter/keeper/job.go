package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mohammadreza-torkaman/wework/x/recruiter/types"
)

// GetJobCount get the total number of job
func (k Keeper) GetJobCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.JobCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetJobCount set the total number of job
func (k Keeper) SetJobCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.JobCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendJob appends a job in the store with a new id and update the count
func (k Keeper) AppendJob(
	ctx sdk.Context,
	job types.Job,
) uint64 {
	// Create the job
	count := k.GetJobCount(ctx)

	// Set the ID of the appended value
	job.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.JobKey))
	appendedValue := k.cdc.MustMarshal(&job)
	store.Set(GetJobIDBytes(job.Id), appendedValue)

	// Update job count
	k.SetJobCount(ctx, count+1)

	return count
}

// SetJob set a specific job in the store
func (k Keeper) SetJob(ctx sdk.Context, job types.Job) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.JobKey))
	b := k.cdc.MustMarshal(&job)
	store.Set(GetJobIDBytes(job.Id), b)
}

// GetJob returns a job from its id
func (k Keeper) GetJob(ctx sdk.Context, id uint64) (val types.Job, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.JobKey))
	b := store.Get(GetJobIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveJob removes a job from the store
func (k Keeper) RemoveJob(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.JobKey))
	store.Delete(GetJobIDBytes(id))
}

// GetAllJob returns all job
func (k Keeper) GetAllJob(ctx sdk.Context) (list []types.Job) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.JobKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Job
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetJobIDBytes returns the byte representation of the ID
func GetJobIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetJobIDFromBytes returns ID in uint64 format from a byte array
func GetJobIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
