package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/phaesoo/citadel/configs"
	"github.com/phaesoo/citadel/db/migrate/migrations"
	"github.com/phaesoo/citadel/pkg/db"
	migrate "github.com/phaesoo/sqlx-migrate"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	mc := configs.Get().Mysql

	flag.Bool("t", false, "to create test db for integration test")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(err)
	}
	test := viper.GetBool("t")

	connString, err := mc.ConnString()
	if err != nil {
		panic(err)
	}
	conn, err := db.NewDB("mysql", connString)
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	// Patch config if test
	if test {
		testDatabase := fmt.Sprintf("%s_test", mc.Database)
		mc.Database = testDatabase

		_, err := conn.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", testDatabase))
		if err != nil {
			log.Fatalf("unable to drop DB `%s`: %s", testDatabase, err.Error())
		}
		_, err = conn.Exec(fmt.Sprintf("CREATE DATABASE %s", testDatabase))
		if err != nil {
			log.Fatalf("unable to create DB `%s`", testDatabase)
		}

		connString, err := mc.ConnString()
		if err != nil {
			panic(err)
		}
		conn, err = db.NewDB("mysql", connString)
		defer conn.Close()
		if err != nil {
			panic(err)
		}
	}

	m := migrate.New(conn, []migrate.Migration{
		migrations.InitTables,
	})
	if err := m.Migrate(); err != nil {
		log.Printf("Migration failed: %s", err.Error())
		if err := m.Rollback(); err != nil {
			log.Printf("Failed to rollback last migration: %s", err.Error())
		}
		log.Printf("Success to rollback last migration")
	}
}
