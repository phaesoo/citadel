package configs

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const (
	fileName  = "config"
	envPrefix = "app"
)

var conf Config

// Config aggregation
type Config struct {
	App        AppConfig
	Mysql      MysqlConfig
	Redis      RedisConfig
	Log        LogConfig
	ElasticAPM ElasticAPMConfig
	Sentry     SentryConfig
}

// Init is explicit initializer for Config
func init() {
	v := initViper()
	conf = Config{
		App:        appConfig(v),
		Mysql:      mysqlConfig(v),
		Redis:      redisConfig(v),
		Log:        logConfig(v),
		ElasticAPM: elasticApmConfig(v),
		Sentry:     sentryConfig(v),
	}
}

// Get returns Config object
func Get() Config {
	return conf
}

func initViper() *viper.Viper {
	v := viper.New()
	v.SetConfigName(fileName)

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	root := "keybox/"
	i := strings.LastIndex(path, root)
	if i != -1 {
		path = path[:i+len(root)]
	}

	log.Print(path)
	v.AddConfigPath(path)

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	v.SetEnvPrefix(envPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// All env vars starts with APP_
	v.AutomaticEnv()
	return v
}
