package checkerstorram

import (
	"checkers-torram/x/checkerstorram/keeper"
	"checkers-torram/x/checkerstorram/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set params
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}

	// Set game count
	k.SetGameCount(ctx, genState.GameCount)

	// Set stored games
	for _, game := range genState.StoredGames {
		k.SetStoredGame(ctx, game)
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get game count
	genesis.GameCount = k.GetGameCount(ctx)

	// Get all games
	store := k.GetStoreService().OpenKVStore(ctx)
	iterator, _ := store.Iterator([]byte(types.GameKey), nil) // nil means iterate to the end
	defer iterator.Close()

	var games []types.StoredGame
	for ; iterator.Valid(); iterator.Next() {
		var game types.StoredGame
		bz := iterator.Value()
		k.GetCdc().MustUnmarshal(bz, &game)
		games = append(games, game)
	}

	genesis.StoredGames = games

	return genesis
}
