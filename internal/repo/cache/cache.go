package cache

import (
	"github.com/phaesoo/keybox/pkg/encrypt"
	"github.com/phaesoo/keybox/pkg/memdb"
)

type Cache struct {
	pool   *memdb.Pool
	cipher *encrypt.AESCipher
}

// NewCache returns implementation of cache layer
func NewCache(pool *memdb.Pool, secretKey string) *Cache {
	cipher, err := encrypt.NewAESCipher([]byte(secretKey))
	if err != nil {
		panic(err)
	}
	return &Cache{pool: pool, cipher: cipher}
}
