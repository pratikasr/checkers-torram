package keeper

import (
	"checkers-torram/x/checkerstorram/types"
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the CheckersTorram service
func NewMsgServerImpl(keeper Keeper) types.CheckersTorramServer {
	return &msgServer{Keeper: keeper}
}

var _ types.CheckersTorramServer = msgServer{}

// CheckersCreateGm creates a new game with timestamps
func (k msgServer) CheckersCreateGm(goCtx context.Context, msg *types.ReqCheckersTorram) (*types.ResCheckersTorram, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get current count and increment
	count := k.GetGameCount(ctx)
	newIndex := strconv.FormatUint(count, 10)

	// Create new game with timestamps
	game := types.StoredGame{
		Index:         newIndex,
		Board:         "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:          "b",
		Black:         msg.Black,
		Red:           msg.Red,
		GameStartTime: ctx.BlockTime().Unix(),
		GameEndTime:   0, // Will be set when game ends
	}

	// Store the game
	k.SetStoredGame(ctx, game)
	k.SetGameCount(ctx, count+1)

	return &types.ResCheckersTorram{
		GameIndex: newIndex,
	}, nil
}
