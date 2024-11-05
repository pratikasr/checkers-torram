package keeper

import (
	"context"

	"checkers-torram/x/checkerstorram/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Params(ctx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	context := sdk.UnwrapSDKContext(ctx)
	params := k.GetParams(context)

	return &types.QueryParamsResponse{Params: params}, nil
}

func (k Keeper) GetGame(ctx context.Context, req *types.QueryGetGameRequest) (*types.QueryGetGameResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	context := sdk.UnwrapSDKContext(ctx)
	game, found := k.GetStoredGame(context, req.Index)
	if !found {
		return nil, status.Error(codes.NotFound, "game not found")
	}

	return &types.QueryGetGameResponse{StoredGame: game}, nil
}
