package repo

import (
	"context"
	"database/sql"

	"github.com/phaesoo/keybox/internal/models"
)

const (
	authKeyTTL = 21600 // 6 hours
)

type authKeyRepo interface {
	AuthKey(ctx context.Context, keyID string) (models.AuthKey, error)
	SetAuthKey(ctx context.Context, authKey models.AuthKey) error
	DeleteAuthKey(ctx context.Context, keyID string) error
}

func (r *repo) AuthKey(ctx context.Context, keyID string) (models.AuthKey, error) {
	var authKey models.AuthKey
	var err error

	authKey, err = r.db.AuthKey(keyID)
	if err != nil {
		if err == sql.ErrNoRows {
			return authKey, ErrNotFound
		}
		return authKey, err
	}
	return authKey, nil
}

func (r *repo) SetAuthKey(ctx context.Context, authKey models.AuthKey) error {
	return r.db.SetAuthKey(authKey)
}

func (r *repo) DeleteAuthKey(ctx context.Context, keyID string) error {
	_, err := r.db.AuthKey(keyID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound
		}
		return err
	}
	return r.db.DeleteAuthKey(keyID)
}
