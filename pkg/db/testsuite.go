package db

import (
	"github.com/phaesoo/citadel/configs"
	"github.com/stretchr/testify/suite"
)

const testDatabase string = "citadel_test"

type TestSuite struct {
	suite.Suite
	Conn *DB
}

func (s *TestSuite) SetupSuite() {
	mc := configs.Get().Mysql
	mc.Database = testDatabase // Set test database

	connString, err := mc.ConnString()
	if err != nil {
		panic(err)
	}

	db, err := NewDB("mysql", connString)
	if err != nil {
		panic(err)
	}
	s.Conn = db
}

func (s *TestSuite) Reset() {
	tx := s.Conn.MustBegin()
	tx.MustExec("TRUNCATE TABLE auth_key;")
	if err := tx.Commit(); err != nil {
		panic(err)
	}
}

func (s *TestSuite) TearDownSuite() {
	s.Reset()
}
