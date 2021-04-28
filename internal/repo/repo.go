package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/phaesoo/keybox/internal/repo/cache"
	"github.com/phaesoo/keybox/internal/repo/db"
	"github.com/phaesoo/keybox/pkg/memdb"
)

type repo struct {
	db    *db.DB
	cache *cache.Cache
}

type Repo interface {
	authKeyRepo
}

// NewRepo returns db implements Repo interface
func NewRepo(conn *sqlx.DB, pool *memdb.Pool, secretKey string) Repo {
	return &repo{
		db:    db.NewDB(conn, secretKey),
		cache: cache.NewCache(pool, secretKey),
	}
}
