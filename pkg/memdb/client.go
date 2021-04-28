package memdb

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"github.com/rafaeljusto/redigomock"
)

type Client = redis.Conn

type Pool = redis.Pool

const (
	// constant db configuration
	maxIdle     int           = 256
	idleTimeout time.Duration = 240 * time.Second

	testDatabase int = 1
)

type Config struct {
	Address      string
	DB           int
	TLSRequired  bool
	AuthRequired bool
	Password     string
	CACert       string
}

// NewPool returns redis client pool
// when a client is required, use `pool.Get()`
// and close it once it is no longer needed.
func NewPool(config Config) *Pool {
	return &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: idleTimeout,
		Dial: func() (Client, error) {
			return Dial(config)
		},
	}
}

func Dial(config Config) (Client, error) {
	client, err := dial(config)
	if err != nil {
		return client, nil
	}
	if config.AuthRequired {
		if _, err := client.Do("AUTH", config.Password); err != nil {
			client.Close()
			return nil, errors.Wrap(err, "auth on redis")
		}
	}
	if config.DB != 0 {
		_, err = client.Do("SELECT", config.DB)
	}
	return client, err
}

// NewMockPool creates a mock pool for testing
func NewMockPool(conn *redigomock.Conn) *Pool {
	return &redis.Pool{
		// Return the same connection mock for each Get() call.
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}
}

// NewTestPool creates a connection pool for the integration test
func NewTestPool(config Config) *Pool {
	config.DB = testDatabase
	return NewPool(config)
}

func dial(config Config) (Client, error) {
	if config.TLSRequired {
		tlsConfig := &tls.Config{}
		if config.CACert != "" {
			certBytes, err := ioutil.ReadFile(config.CACert)
			if err != nil {
				return nil, err
			}

			caCertPool := x509.NewCertPool()
			ok := caCertPool.AppendCertsFromPEM(certBytes)
			if !ok {
				panic("failed to parse ca cert")
			}
			tlsConfig.RootCAs = caCertPool
		}

		return redis.Dial(
			"tcp",
			config.Address,
			redis.DialConnectTimeout(2*time.Second),
			redis.DialTLSConfig(tlsConfig),
			redis.DialUseTLS(true),
			redis.DialTLSSkipVerify(true),
		)
	}
	return redis.Dial(
		"tcp",
		config.Address,
	)
}
