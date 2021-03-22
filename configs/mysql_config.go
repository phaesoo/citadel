package configs

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type MysqlConfig struct {
	Host        string
	Port        int
	Database    string
	User        string
	Password    string
	TLSRequired bool
	// Path to CA File for SSL connection.
	CA      string
	Logging bool
}

func mysqlConfig(v *viper.Viper) MysqlConfig {
	return MysqlConfig{
		Host:        v.GetString("mysql.host"),
		Port:        v.GetInt("mysql.port"),
		Database:    v.GetString("mysql.database"),
		User:        v.GetString("mysql.user"),
		Password:    v.GetString("mysql.password"),
		TLSRequired: v.GetBool("mysql.tlsRequired"),
		CA:          v.GetString("mysql.ca"),
		Logging:     v.GetBool("mysql.logging"),
	}
}

func (c *MysqlConfig) ConnString() (string, error) {
	connStr := fmt.Sprintf("%s:%s@(%s:%d)/%s", c.User, c.Password, c.Host, c.Port, c.Database)
	if c.TLSRequired {
		certBytes, err := ioutil.ReadFile(c.CA)
		if err != nil {
			return "", err
		}
		caCertPool := x509.NewCertPool()
		if ok := caCertPool.AppendCertsFromPEM(certBytes); !ok {
			return "", errors.Errorf("Fail to parse ca")
		}
		tlsConfig := &tls.Config{
			InsecureSkipVerify: false,
			RootCAs:            caCertPool,
		}
		tlsName := "custom"
		mysql.RegisterTLSConfig(tlsName, tlsConfig)
		connStr = fmt.Sprintf("%s?tls=%s", connStr, tlsName)
	}
	return connStr, nil
}
