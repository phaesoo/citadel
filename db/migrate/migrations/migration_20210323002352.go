package migrations

import (
	"github.com/jmoiron/sqlx"
	migrate "github.com/phaesoo/sqlx-migrate"
)

var InitTables = migrate.Migration{
	ID:   "20210323002352",
	Name: "InitTables",
	Migrate: func(tx *sqlx.Tx) error {
		return nil
	},
	Rollback: func(tx *sqlx.Tx) error {
		return nil
	},
}
