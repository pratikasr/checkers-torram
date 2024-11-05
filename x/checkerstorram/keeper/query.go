// x/checkerstorram/keeper/query.go
package keeper

import (
	"checkers-torram/x/checkerstorram/types"
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type queryServer struct {
	Keeper
}

// NewQueryServerImpl returns an implementation of the QueryServer interface
func NewQueryServerImpl(k Keeper) types.QueryServer {
	return &queryServer{Keeper: k}
}

// GetStoredGame implements types.QueryServer
func (q queryServer) GetStoredGame(goCtx context.Context, req *types.QueryGetStoredGameRequest) (*types.QueryGetStoredGameResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	game, found := q.Keeper.GetStoredGame(ctx, req.Index)
	if !found {
		return nil, status.Error(codes.NotFound, "game not found")
	}

	return &types.QueryGetStoredGameResponse{
		StoredGame: game,
	}, nil
}
