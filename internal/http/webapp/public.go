package webapp

import (
	"github.com/feditools/relay/internal/http/template"
	"net/http"
)

func makePublicNavbar(_ *http.Request) template.Navbar {
	return template.Navbar{}
}
