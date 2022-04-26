package config

// Values contains the type of each value.
type Values struct {
	ConfigPath string
	LogLevel   string

	// application
	ApplicationName string
	SoftwareVersion string

	// database
	DbType          string
	DbAddress       string
	DbPort          int
	DbUser          string
	DbPassword      string
	DbDatabase      string
	DbTLSMode       string
	DbTLSCACert     string
	DbLoadTestData  bool
	DbEncryptionKey string

	// server
	ServerExternalHostname string
	ServerHTTPBind         string
	ServerMinifyHTML       bool

	// metrics
	MetricsStatsDAddress string
	MetricsStatsDPrefix  string
}

// Defaults contains the default values
var Defaults = Values{
	ConfigPath: "",
	LogLevel:   "info",

	// application
	ApplicationName: "relay",

	// database
	DbType:         "postgres",
	DbAddress:      "",
	DbPort:         5432,
	DbUser:         "",
	DbPassword:     "",
	DbDatabase:     "relay",
	DbTLSMode:      "disable",
	DbTLSCACert:    "",
	DbLoadTestData: false,

	// server
	ServerExternalHostname: "localhost",
	ServerHTTPBind:         ":5000",
	ServerMinifyHTML:       true,

	// metrics
	MetricsStatsDAddress: "localhost:8125",
	MetricsStatsDPrefix:  "relay",
}
