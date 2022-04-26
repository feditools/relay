package config

// KeyNames is a struct that contains the names of keys.
type KeyNames struct {
	LogLevel   string
	ConfigPath string

	// application
	ApplicationName string
	SoftwareVersion string
	TokenSalt       string

	// database
	DbType          string
	DbAddress       string
	DbPort          string
	DbUser          string
	DbPassword      string
	DbDatabase      string
	DbTLSMode       string
	DbTLSCACert     string
	DbLoadTestData  string
	DbEncryptionKey string

	// server
	ServerExternalHostname string
	ServerHTTPBind         string
	ServerMinifyHTML       string

	// metrics
	MetricsStatsDAddress string
	MetricsStatsDPrefix  string
}

// Keys contains the names of config keys.
var Keys = KeyNames{
	ConfigPath: "config-path", // CLI only
	LogLevel:   "log-level",

	// application
	ApplicationName: "application-name",
	SoftwareVersion: "software-version", // Set at build
	TokenSalt:       "token-salt",

	// database
	DbType:          "db-type",
	DbAddress:       "db-address",
	DbPort:          "db-port",
	DbUser:          "db-user",
	DbPassword:      "db-password",
	DbDatabase:      "db-database",
	DbTLSMode:       "db-tls-mode",
	DbTLSCACert:     "db-tls-ca-cert",
	DbLoadTestData:  "test-data", // CLI only
	DbEncryptionKey: "db-crypto-key",

	// server
	ServerExternalHostname: "external-hostname",
	ServerHTTPBind:         "http-bind",
	ServerMinifyHTML:       "minify-html",

	// metrics
	MetricsStatsDAddress: "statsd-addr",
	MetricsStatsDPrefix:  "statsd-prefix",
}
