package cache

import (
	"github.com/gomodule/redigo/redis"
	"github.com/phaesoo/keybox/configs"
	"github.com/phaesoo/keybox/pkg/memdb"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	Pool *redis.Pool
	Conn redis.Conn
}

func (s *TestSuite) ClearKeys(keys ...string) {
	for _, key := range keys {
		_, err := s.Conn.Do("DEL", key)
		s.NoError(err)
	}
}

func (s *TestSuite) SetupSuite() {
	rc := configs.Get().Redis
	s.Pool = memdb.NewTestPool(
		memdb.Config{
			Address:      rc.Address(),
			DB:           1,
			TLSRequired:  rc.TLSRequired,
			AuthRequired: rc.AuthRequired,
			Password:     rc.Password,
			CACert:       rc.CACert,
		},
	)
	s.Conn = s.Pool.Get()
}

func (s *TestSuite) TearDownSuite() {
	s.NoError(s.Conn.Flush())
	s.NoError(s.Conn.Close())
	s.NoError(s.Pool.Close())
}
