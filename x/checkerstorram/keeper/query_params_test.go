package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "checkers-torram/testutil/keeper"
	"checkers-torram/x/checkerstorram/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.CheckerstorramKeeper(t)
	params := types.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
