package types

const (
    // ModuleName defines the name of the module
    ModuleName = "encryption"

    // StoreKey is the primary module store key for encryption
    StoreKey = ModuleName

    // RouterKey is the message route for encryption-related transactions
    RouterKey = ModuleName

    // QuerierRoute defines the module’s query route for encryption data
    QuerierRoute = ModuleName

    // MemStoreKey is the in-memory store key for the encryption module
    MemStoreKey = "mem_encryption"
)

var (
    // KeyPrefix is a common prefix for encryption module keys in the KVStore
    KeyPrefix = []byte{0x01}
)

// KeyAES is a unique key for storing AES encryption data in the KVStore
func KeyAES() []byte {
    return append(KeyPrefix, []byte("AES")...)
}

// KeyIP is a unique key for storing IP-based encryption data in the KVStore
func KeyIP() []byte {
    return append(KeyPrefix, []byte("IP")...)
}
