package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

// AppConfig is config struct for app
type AppConfig struct {
	Name    string
	Host    string
	Port    int
	Profile bool
	Metrics bool
}

func appConfig(v *viper.Viper) AppConfig {
	return AppConfig{
		Name:    v.GetString("citadel.name"),
		Host:    v.GetString("citadel.host"),
		Port:    v.GetInt("citadel.port"),
		Profile: v.GetBool("citadel.profile"),
		Metrics: v.GetBool("citadel.metrics"),
	}
}

func (c *AppConfig) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
