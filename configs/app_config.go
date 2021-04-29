package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

// AppConfig is config struct for app
type AppConfig struct {
	Name      string
	Host      string
	Port      int
	SecretKey string
	Profile   bool
	Metrics   bool
}

func appConfig(v *viper.Viper) AppConfig {
	return AppConfig{
		Name:      v.GetString("keybox.name"),
		Host:      v.GetString("keybox.host"),
		Port:      v.GetInt("keybox.port"),
		SecretKey: v.GetString("keybox.secretKey"),
		Profile:   v.GetBool("keybox.profile"),
		Metrics:   v.GetBool("keybox.metrics"),
	}
}

func (c *AppConfig) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
