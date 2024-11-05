package types

const (
	// ModuleName defines the module name
	ModuleName = "checkerstorram"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_checkerstorram"

	// Game-related keys
	GameKey      = "Game/value/"
	GameCountKey = "Game/count/"
)

var (
	ParamsKey = []byte("p_checkerstorram")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// GetGameIDBytes returns the byte representation of the gameID
func GetGameIDBytes(gameID string) []byte {
	return []byte(GameKey + gameID)
}

// GetGameIDFromBytes returns gameID in string form from a byte array
func GetGameIDFromBytes(bz []byte) string {
	return string(bz[len(GameKey):])
}
