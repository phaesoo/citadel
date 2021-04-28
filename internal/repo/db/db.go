package db

import (
	"github.com/jmoiron/sqlx"
)

type DB struct {
	conn      *sqlx.DB
	secretKey string
}

// NewRepo returns db implements Repo interface
func NewDB(conn *sqlx.DB, secretKey string) *DB {
	return &DB{conn: conn, secretKey: secretKey}
}
