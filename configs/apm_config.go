package configs

import (
	"github.com/spf13/viper"
)

// ElasticAPMConfig is a set of config options for Elastic APM. See the APM middleware and [Elastic
// APM Docs](https://www.elastic.co/guide/en/apm/agent/go/current/configuration.html) for more
// information.
type ElasticAPMConfig struct {
	ServerURL      string
	ServiceName    string
	ServiceVersion string
	RedactedFields string
	Environment    string
	Active         bool
	SecretToken    string
	LogFile        string
	LogLevel       string
}

const (
	serverURLVar      = "apm.serverUrl"
	serviceNameVar    = "apm.serviceName"
	serviceVersionVar = "apm.serviceVersion"
	activeVar         = "apm.active"
	redactedFieldsVar = "apm.redactedFields"
	environmentVar    = "apm.environment"
	secretTokenVar    = "apm.secretToken"
	logFileVar        = "apm.logFile"
	logLevelVar       = "apm.logLevel"
)

func elasticApmConfig(v *viper.Viper) ElasticAPMConfig {
	// BindEnv only returns an error if no proper string input is provided
	_ = v.BindEnv(serverURLVar, "ELASTIC_APM_SERVER_URL")
	_ = v.BindEnv(serviceNameVar, "ELASTIC_APM_SERVICE_NAME")
	_ = v.BindEnv(serviceVersionVar, "ELASTIC_APM_SERVICE_VERSION")
	_ = v.BindEnv(redactedFieldsVar, "ELASTIC_APM_SANITIZE_FIELD_NAMES")
	_ = v.BindEnv(environmentVar, "ELASTIC_APM_ENVIRONMENT")
	_ = v.BindEnv(activeVar, "ELASTIC_APM_ACTIVE")
	_ = v.BindEnv(secretTokenVar, "ELASTIC_APM_SECRET_TOKEN")
	_ = v.BindEnv(logFileVar, "ELASTIC_APM_LOG_FILE")
	_ = v.BindEnv(logLevelVar, "ELASTIC_APM_LOG_LEVEL")

	cfg := ElasticAPMConfig{
		ServerURL:      v.GetString(serverURLVar),
		ServiceName:    v.GetString(serviceNameVar),
		ServiceVersion: v.GetString(serviceVersionVar),
		RedactedFields: v.GetString(redactedFieldsVar),
		Environment:    v.GetString(environmentVar),
		Active:         v.GetBool(activeVar),
		SecretToken:    v.GetString(secretTokenVar),
		LogFile:        v.GetString(logFileVar),
		LogLevel:       v.GetString(logLevelVar),
	}
	return cfg
}
