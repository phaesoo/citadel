package db

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const (
	retryCount = 1
)

type DB = sqlx.DB

// NewDB returns connected Client
func NewDB(driverName, connString string) (*DB, error) {
	var conn *sqlx.DB
	var err error
	for i := 0; i < retryCount; i++ {
		conn, err = sqlx.Connect(driverName, connString)
		if err != nil {
			log.Print(err)
			time.Sleep(time.Second)
			continue
		}
		return conn, nil
	}
	return nil, errors.Wrap(err, "DB connect")
}
