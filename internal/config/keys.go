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
	EncryptionKey       string
	SoftwareVersion     string

	// database
	DbType      string
	DbAddress   string
	DbPort      string
	DbUser      string
	DbPassword  string
	DbDatabase  string
	DbTLSMode   string
	DbTLSCACert string

	// redis
	RedisAddress  string
	RedisDB       string
	RedisPassword string

	// runner
	RunnerConcurrency string

	// server
	ServerExternalHostname string
	ServerHTTPBind         string
	ServerMinifyHTML       string
	ServerRoles            string

	// webapp
	WebappBootstrapCSSURI         string
	WebappBootstrapCSSIntegrity   string
	WebappBootstrapJSURI          string
	WebappBootstrapJSIntegrity    string
	WebappFontAwesomeCSSURI       string
	WebappFontAwesomeCSSIntegrity string
	WebappLogoSrcDark             string
	WebappLogoSrcLight            string

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
	EncryptionKey:       "encryption-key",
	SoftwareVersion:     "software-version", // Set at build

	// database
	DbType:      "db-type",
	DbAddress:   "db-address",
	DbPort:      "db-port",
	DbUser:      "db-user",
	DbPassword:  "db-password",
	DbDatabase:  "db-database",
	DbTLSMode:   "db-tls-mode",
	DbTLSCACert: "db-tls-ca-cert",

	// redis
	RedisAddress:  "redis-address",
	RedisDB:       "redis-db",
	RedisPassword: "redis-password",

	// runner
	RunnerConcurrency: "runner-concurrency",

	// server
	ServerExternalHostname: "external-hostname",
	ServerHTTPBind:         "http-bind",
	ServerMinifyHTML:       "minify-html",
	ServerRoles:            "server-role",

	// webapp
	WebappBootstrapCSSURI:         "webapp-bootstrap-css-uri",
	WebappBootstrapCSSIntegrity:   "webapp-bootstrap-css-integrity",
	WebappBootstrapJSURI:          "webapp-bootstrap-js-uri",
	WebappBootstrapJSIntegrity:    "webapp-bootstrap-js-integrity",
	WebappFontAwesomeCSSURI:       "webapp-fontawesome-css-uri",
	WebappFontAwesomeCSSIntegrity: "webapp-fontawesome-css-integrity",
	WebappLogoSrcDark:             "webapp-logo-src-dark",
	WebappLogoSrcLight:            "webapp-logo-src-light",

	// metrics
	MetricsStatsDAddress: "statsd-addr",
	MetricsStatsDPrefix:  "statsd-prefix",
}
