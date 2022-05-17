package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/mohammadreza-torkaman/wework/x/recruiter/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) JobAll(c context.Context, req *types.QueryAllJobRequest) (*types.QueryAllJobResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var jobs []types.Job
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	jobStore := prefix.NewStore(store, types.KeyPrefix(types.JobKey))

	pageRes, err := query.Paginate(jobStore, req.Pagination, func(key []byte, value []byte) error {
		var job types.Job
		if err := k.cdc.Unmarshal(value, &job); err != nil {
			return err
		}

		jobs = append(jobs, job)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllJobResponse{Job: jobs, Pagination: pageRes}, nil
}

func (k Keeper) Job(c context.Context, req *types.QueryGetJobRequest) (*types.QueryGetJobResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	job, found := k.GetJob(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetJobResponse{Job: job}, nil
}
