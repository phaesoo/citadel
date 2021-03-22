package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type RedisConfig struct {
	Host         string
	Port         string
	Database     int
	TLSRequired  bool
	AuthRequired bool
	Password     string `json:"-"`
	CACert       string `json:"-"`
	Logging      bool
}

func redisConfig(v *viper.Viper) RedisConfig {
	return RedisConfig{
		Host:         v.GetString("redis.host"),
		Port:         v.GetString("redis.port"),
		Database:     v.GetInt("redis.database"),
		TLSRequired:  v.GetBool("redis.tlsRequired"),
		AuthRequired: v.GetBool("redis.authRequired"),
		Password:     v.GetString("redis.password"),
		CACert:       v.GetString("redis.ca"),
		Logging:      v.GetBool("redis.logging"),
	}
}

func (c RedisConfig) Address() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
