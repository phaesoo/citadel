package cache

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/phaesoo/keybox/internal/models"
	"github.com/pkg/errors"
)

const (
	authKeyHashPrefix = "keybox:auth-key:%s:%s"
)

func generateKey(userID, keyID string) string {
	return fmt.Sprintf(authKeyHashPrefix, userID, keyID)
}

func (c *Cache) AuthKey(userID, keyID string) (models.AuthKey, error) {
	conn := c.pool.Get()
	defer conn.Close()

	key := generateKey(userID, keyID)

	val, err := redis.Values(conn.Do("HGETALL", key))
	if err != nil {
		return models.AuthKey{}, errors.Wrap(err, "Get auth key from redis")
	}

	var authKey models.AuthKey
	if err := redis.ScanStruct(val, &authKey); err != nil {
		return authKey, err
	} else if authKey == (models.AuthKey{}) {
		return authKey, ErrNotFound
	}

	authKey.PublicKey, err = c.cipher.Decrypt(authKey.PublicKey)
	authKey.PrivateKey, err = c.cipher.Decrypt(authKey.PrivateKey)
	if err != nil {
		return authKey, err
	}

	return authKey, nil
}

func (c *Cache) SetAuthKey(authKey models.AuthKey, ttl int) error {
	conn := c.pool.Get()
	defer conn.Close()

	key := generateKey(authKey.UserID, authKey.KeyID)

	var err error
	authKey.PublicKey, err = c.cipher.Encrypt(authKey.PublicKey)
	authKey.PrivateKey, err = c.cipher.Encrypt(authKey.PrivateKey)
	if err != nil {
		return err
	}

	_, err = conn.Do("HMSET", redis.Args{key}.AddFlat(authKey)...)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, ttl)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cache) DeleteAuthKey(userID, keyID string) error {
	conn := c.pool.Get()
	defer conn.Close()

	key := generateKey(userID, keyID)

	_, err := conn.Do("DEL", key)
	if err != nil {
		return err
	}

	return nil
}
