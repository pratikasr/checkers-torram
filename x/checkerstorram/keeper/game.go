// keeper/game.go
package keeper

import (
	"checkers-torram/x/checkerstorram/types"
	"encoding/binary"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetGameCount sets the total number of games
func (k Keeper) SetGameCount(ctx sdk.Context, count uint64) {
	store := k.storeService.OpenKVStore(ctx)
	countBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(countBytes, count)
	store.Set([]byte(types.GameCountKey), countBytes)
}

// GetGameCount gets the total number of games
func (k Keeper) GetGameCount(ctx sdk.Context) uint64 {
	store := k.storeService.OpenKVStore(ctx)
	bz, _ := store.Get([]byte(types.GameCountKey))
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

// SetStoredGame sets a specific game in store
func (k Keeper) SetStoredGame(ctx sdk.Context, game types.StoredGame) {
	store := k.storeService.OpenKVStore(ctx)
	gameBytes := k.cdc.MustMarshal(&game)
	store.Set(types.GetGameIDBytes(game.Index), gameBytes)
}

// GetStoredGame gets a specific game from the store
func (k Keeper) GetStoredGame(ctx sdk.Context, gameID string) (types.StoredGame, bool) {
	store := k.storeService.OpenKVStore(ctx)
	bz, _ := store.Get(types.GetGameIDBytes(gameID))
	if bz == nil {
		return types.StoredGame{}, false
	}
	var game types.StoredGame
	k.cdc.MustUnmarshal(bz, &game)
	return game, true
}
