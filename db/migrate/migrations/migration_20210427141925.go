package migrations

import (
	"github.com/jmoiron/sqlx"
	migrate "github.com/phaesoo/sqlx-migrate"
)

var InitTables = migrate.Migration{
	ID:   "20210427141925",
	Name: "Init DB",
	Migrate: func(tx *sqlx.Tx) error {
		if _, err := tx.Exec(`
		CREATE TABLE auth_key (
			id INT PRIMARY KEY AUTO_INCREMENT,
			key_id VARCHAR(255) UNIQUE,
			public_pem VARCHAR(3000),
			private_pem VARCHAR(3000),
			user_id VARCHAR(255)
		) ENGINE=INNODB;
		`); err != nil {
			return err
		}
		return nil
	},
	Rollback: func(tx *sqlx.Tx) error {
		return nil
	},
}
