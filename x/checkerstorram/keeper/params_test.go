package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "checkers-torram/testutil/keeper"
	"checkers-torram/x/checkerstorram/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.CheckerstorramKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
