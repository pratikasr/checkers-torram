package keeper_test

import (
	keepertest "checkers-torram/testutil/keeper"
	"checkers-torram/x/checkerstorram/keeper"
	"checkers-torram/x/checkerstorram/types"
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

// setupMsgServer returns a keeper, a message server, and a context
func setupMsgServer(t testing.TB) (types.CheckersTorramServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.CheckerstorramKeeper(t)
	return keeper.NewMsgServerImpl(k), k, sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, k, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}

func TestCheckersCreateGm(t *testing.T) {
	msgServer, keeper, context := setupMsgServer(t)

	tests := []struct {
		desc    string
		request *types.ReqCheckersTorram
		err     bool
	}{
		{
			desc: "Create First Game Success",
			request: &types.ReqCheckersTorram{
				Creator: "cosmos1def",
				Black:   "cosmos1def",
				Red:     "cosmos1xyz",
			},
			err: false,
		},
		{
			desc: "Create Second Game Success",
			request: &types.ReqCheckersTorram{
				Creator: "cosmos1def",
				Black:   "cosmos1abc",
				Red:     "cosmos1xyz",
			},
			err: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := msgServer.CheckersCreateGm(context, tc.request)

			if tc.err {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, response)

			// Verify game was stored
			game, found := keeper.GetStoredGame(sdk.UnwrapSDKContext(context), response.GameIndex)
			require.True(t, found)

			// Verify game fields
			require.Equal(t, tc.request.Black, game.Black)
			require.Equal(t, tc.request.Red, game.Red)
			require.Equal(t, "b", game.Turn)
			require.Equal(t, "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*", game.Board)

			// Verify timestamps
			ctx := sdk.UnwrapSDKContext(context)
			require.Equal(t, ctx.BlockTime().Unix(), game.GameStartTime)
			require.Equal(t, int64(0), game.GameEndTime)

			// Verify game count increased
			count := keeper.GetGameCount(sdk.UnwrapSDKContext(context))
			expectedIndex, err := strconv.ParseUint(response.GameIndex, 10, 64)
			require.NoError(t, err)
			require.Equal(t, expectedIndex+1, count)
		})
	}
}

func TestCheckersCreateGm_GameCount(t *testing.T) {
	msgServer, keeper, context := setupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(context)

	// Create multiple games
	for i := 0; i < 3; i++ {
		request := &types.ReqCheckersTorram{
			Creator: "cosmos1def",
			Black:   "cosmos1def",
			Red:     "cosmos1xyz",
		}
		response, err := msgServer.CheckersCreateGm(context, request)
		require.NoError(t, err)
		require.NotNil(t, response)

		// Verify game index matches count
		expectedIndex := strconv.FormatUint(uint64(i), 10)
		require.Equal(t, expectedIndex, response.GameIndex)
	}

	// Verify final count
	count := keeper.GetGameCount(ctx)
	require.Equal(t, uint64(3), count)
}
