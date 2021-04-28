package db

import (
	"fmt"

	"github.com/phaesoo/keybox/internal/models"
	rdb "github.com/phaesoo/keybox/pkg/db"
)

func (db *DB) AuthKey(userID, keyID string) (models.AuthKey, error) {
	k := struct {
		ID         int    `db:"id"`
		KeyID      string `db:"key_id"`
		PublicKey  string `db:"public_key"`
		PrivateKey string `db:"private_key"`
		UserID     string `db:"user_id"`
	}{}

	if err := db.conn.Get(&k, fmt.Sprintf(`
	SELECT 
		id,
		key_id,
		AES_DECRYPT(UNHEX(public_key), '%s') as public_key,
		AES_DECRYPT(UNHEX(private_key), '%s') as private_key,
		user_id
	FROM auth_key
	WHERE user_id = ? and key_id = ?
	`, db.secretKey, db.secretKey), userID, keyID); err != nil {
		return models.AuthKey{}, err
	}

	return models.AuthKey{
		ID:         k.ID,
		KeyID:      k.KeyID,
		PublicKey:  k.PublicKey,
		PrivateKey: k.PrivateKey,
		UserID:     k.UserID,
	}, nil
}

func (db *DB) SetAuthKey(authKey models.AuthKey) error {
	return rdb.WithTransaction(db.conn, func(tx rdb.Transaction) error {
		_, err := tx.Exec(fmt.Sprintf(`
		INSERT INTO auth_key (key_id, public_key, private_key, user_id)
		VALUES (?, HEX(AES_ENCRYPT(?, '%s')), HEX(AES_ENCRYPT(?, '%s')), ?, ?)
		`, db.secretKey, db.secretKey), authKey.KeyID, authKey.PublicKey, authKey.PrivateKey, authKey.UserID)
		return err
	})
}

func (db *DB) DeleteAuthKey(userID, keyID string) error {
	_, err := db.conn.Exec(`
	DELETE FROM auth_key
	WHERE user_id = ? and key_id = ?
	`, userID, keyID)
	return err
}
