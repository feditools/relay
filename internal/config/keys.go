package config

// KeyNames is a struct that contains the names of keys.
type KeyNames struct {
	LogLevel   string
	ConfigPath string

	// application
	ActorKeySize        string
	ApplicationName     string
	CachedActivityLimit string
	CachedActorLimit    string
	CachedDigestLimit   string
	SoftwareVersion     string

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

	// running
	RunnerConcurrency string

	// server
	ServerExternalHostname string
	ServerHTTPBind         string
	ServerMinifyHTML       string
	ServerRoles            string

	// metrics
	MetricsStatsDAddress string
	MetricsStatsDPrefix  string
}

// Keys contains the names of config keys.
var Keys = KeyNames{
	ConfigPath: "config-path", // CLI only
	LogLevel:   "log-level",

	// application
	ActorKeySize:        "actor-key-size",
	ApplicationName:     "application-name",
	CachedActivityLimit: "cached-activity-limit",
	CachedActorLimit:    "cached-actor-limit",
	CachedDigestLimit:   "cached-digest-limit",
	SoftwareVersion:     "software-version", // Set at build

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

	// runner
	RunnerConcurrency: "runner-concurrency",

	// server
	ServerExternalHostname: "external-hostname",
	ServerHTTPBind:         "http-bind",
	ServerMinifyHTML:       "minify-html",
	ServerRoles:            "server-role",

	// metrics
	MetricsStatsDAddress: "statsd-addr",
	MetricsStatsDPrefix:  "statsd-prefix",
}
