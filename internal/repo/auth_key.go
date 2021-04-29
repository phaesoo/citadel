package repo

import (
	"context"
	"database/sql"

	"github.com/phaesoo/keybox/internal/models"
	"github.com/phaesoo/keybox/internal/repo/cache"
)

const (
	cacheTTL = 3600 // 1 hour
)

type authKeyRepo interface {
	AuthKey(ctx context.Context, userID, keyID string) (models.AuthKey, error)
	SetAuthKey(ctx context.Context, authKey models.AuthKey) error
	DeleteAuthKey(ctx context.Context, userID, keyID string) error
}

func (r *repo) AuthKey(ctx context.Context, userID, keyID string) (models.AuthKey, error) {
	var authKey models.AuthKey
	var err error
	authKey, err = r.cache.AuthKey(userID, keyID)
	if err != nil {
		if err != cache.ErrNotFound {
			return authKey, err
		}

		authKey, err = r.db.AuthKey(userID, keyID)
		if err != nil {
			if err == sql.ErrNoRows {
				return authKey, ErrNotFound
			}
			return authKey, err
		}

		if err := r.cache.SetAuthKey(authKey, cacheTTL); err != nil {
			return authKey, err
		}
	}
	return authKey, nil
}

func (r *repo) SetAuthKey(ctx context.Context, authKey models.AuthKey) error {
	return r.db.SetAuthKey(authKey)
}

func (r *repo) DeleteAuthKey(ctx context.Context, userID, keyID string) error {
	_, err := r.db.AuthKey(userID, keyID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound
		}
		return err
	}
	if err := r.cache.DeleteAuthKey(userID, keyID); err != nil {
		return err
	}
	return r.db.DeleteAuthKey(userID, keyID)
}
