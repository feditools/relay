package config

// Values contains the type of each value.
type Values struct {
	ConfigPath string
	LogLevel   string

	// application
	ActorKeySize        int
	ApplicationName     string
	CachedActivityLimit int
	CachedActorLimit    int
	CachedDigestLimit   int
	EncryptionKey       string
	SoftwareVersion     string

	// database
	DbType      string
	DbAddress   string
	DbPort      int
	DbUser      string
	DbPassword  string
	DbDatabase  string
	DbTLSMode   string
	DbTLSCACert string

	// redis
	RedisAddress  string
	RedisDB       int
	RedisPassword string

	// runner
	RunnerConcurrency int

	// server
	ServerExternalHostname string
	ServerHTTPBind         string
	ServerMinifyHTML       bool
	ServerRoles            []string

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

// Defaults contains the default values
var Defaults = Values{
	ConfigPath: "",
	LogLevel:   "info",

	// application
	ActorKeySize:        2048,
	ApplicationName:     "feditools-relay",
	CachedActivityLimit: 1024,
	CachedActorLimit:    1024,
	CachedDigestLimit:   1024,

	// database
	DbType:      "postgres",
	DbAddress:   "",
	DbPort:      5432,
	DbUser:      "",
	DbPassword:  "",
	DbDatabase:  "relay",
	DbTLSMode:   "disable",
	DbTLSCACert: "",

	// redis
	RedisAddress: "localhost:6379",
	RedisDB:      0,

	// runner
	RunnerConcurrency: 4,

	// server
	ServerExternalHostname: "localhost",
	ServerHTTPBind:         ":5000",
	ServerMinifyHTML:       true,
	ServerRoles: []string{
		ServerRoleActivityPub,
		ServerRoleWebapp,
	},

	// webapp
	WebappBootstrapCSSURI:         "https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css",
	WebappBootstrapCSSIntegrity:   "sha384-1BmE4kWBq78iYhFldvKuhfTAU6auU8tT94WrHftjDbrCEXSU1oBoqyl2QvZ6jIW3",
	WebappBootstrapJSURI:          "https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/js/bootstrap.bundle.min.js",
	WebappBootstrapJSIntegrity:    "sha384-ka7Sk0Gln4gmtz2MlQnikT1wXgYsOg+OMhuP+IlRH9sENBO0LRn5q+8nbTov4+1p",
	WebappFontAwesomeCSSURI:       "https://cdn.fedi.tools/vendor/fontawesome-free-6.1.1/css/all.min.css",
	WebappFontAwesomeCSSIntegrity: "sha384-/frq1SRXYH/bSyou/HUp/hib7RVN1TawQYja658FEOodR/FQBKVqT9Ol+Oz3Olq5",
	WebappLogoSrcDark:             "https://cdn.fedi.tools/img/feditools-logo-dark.svg",
	WebappLogoSrcLight:            "https://cdn.fedi.tools/img/feditools-logo-light.svg",

	// metrics
	MetricsStatsDAddress: "localhost:8125",
	MetricsStatsDPrefix:  "relay",
}
