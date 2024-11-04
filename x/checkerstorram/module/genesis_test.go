package checkerstorram_test

import (
	"testing"

	keepertest "checkers-torram/testutil/keeper"
	"checkers-torram/testutil/nullify"
	checkerstorram "checkers-torram/x/checkerstorram/module"
	"checkers-torram/x/checkerstorram/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CheckerstorramKeeper(t)
	checkerstorram.InitGenesis(ctx, k, genesisState)
	got := checkerstorram.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
