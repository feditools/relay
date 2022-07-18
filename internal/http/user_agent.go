package http

import (
	"fmt"
	"github.com/feditools/relay/internal/config"
	"github.com/spf13/viper"
	"sync"
)

var (
	initOnce      sync.Once
	userAgentLock sync.RWMutex

	userAgent string
)

// doInit sets the User-Agent for all subsequent requests
func doInit() {
	userAgentLock.Lock()
	userAgent = fmt.Sprintf("Go-http-client/2.0 (%s/%s; +https://%s/)",
		viper.GetString(config.Keys.ApplicationName),
		viper.GetString(config.Keys.SoftwareVersion),
		viper.GetString(config.Keys.ServerExternalHostname),
	)
	userAgentLock.Unlock()
}

// GetUserAgent returns the generated http User-Agent
func GetUserAgent() string {
	initOnce.Do(doInit)

	userAgentLock.RLock()
	ua := userAgent
	userAgentLock.RUnlock()
	return ua
}
