package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// BlockFlocGin is a gin middleware that applies a header to tell Google Chrome to disable cohort tracking
func BlockFlocGin(c *gin.Context) {
	c.Header("Permissions-Policy", "interest-cohort=()")
}

// BlockFlocMux is a mux middleware that applies a header to tell Google Chrome to disable cohort tracking
func BlockFlocMux(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Permissions-Policy", "interest-cohort=()")
		next.ServeHTTP(w, r)
	})
}
