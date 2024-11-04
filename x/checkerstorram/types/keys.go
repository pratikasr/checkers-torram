package types

const (
	// ModuleName defines the module name
	ModuleName = "checkerstorram"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_checkerstorram"
)

var (
	ParamsKey = []byte("p_checkerstorram")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
