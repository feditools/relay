package webapp

import (
	"fmt"
	libtemplate "github.com/feditools/go-lib/template"
	"github.com/feditools/relay/internal/http/template"
	"github.com/feditools/relay/internal/path"
	"net/http"
	"strings"
)

func (m *Module) returnErrorPage(w http.ResponseWriter, r *http.Request, code int, errStr string) {
	l := logger.WithField("func", "returnErrorPage")

	// Init template variables
	tmplVars := &template.Error{}
	err := m.initTemplate(w, r, tmplVars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	// add error css file
	signature, err := getSignature(strings.TrimPrefix(path.FileErrorCSS, "/"))
	if err != nil {
		l.Errorf("getting signature for %s: %s", path.FileErrorCSS, err.Error())
	}
	tmplVars.AddHeadLink(libtemplate.HeadLink{
		HRef:        path.FileErrorCSS,
		Rel:         "stylesheet",
		CrossOrigin: "anonymous",
		Integrity:   signature,
	})

	// set image
	tmplVars.Image = m.logoSrcDark

	// set text
	tmplVars.Header = fmt.Sprintf("%d", code)
	tmplVars.SubHeader = http.StatusText(code)
	tmplVars.PageTitle = fmt.Sprintf("%d - %s", code, http.StatusText(code))
	tmplVars.Paragraph = errStr

	// set top button
	switch code {
	case http.StatusUnauthorized:
		tmplVars.ButtonHRef = path.AppLogin
		tmplVars.ButtonLabel = "Login"
	default:
		tmplVars.ButtonHRef = path.AppHome
		tmplVars.ButtonLabel = "Home"
	}

	w.WriteHeader(code)
	err = m.executeTemplate(w, "error", tmplVars)
	if err != nil {
		logger.Errorf("could not render error template: %s", err.Error())
	}
}

func (m *Module) methodNotAllowedHandler() http.Handler {
	// wrap in middleware since middlware isn't run on error pages
	return m.WrapInMiddlewares(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.returnErrorPage(w, r, http.StatusMethodNotAllowed, r.Method)
	}))
}

func (m *Module) notFoundHandler() http.Handler {
	// wrap in middleware since middlware isn't run on error pages
	return m.WrapInMiddlewares(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.returnErrorPage(w, r, http.StatusNotFound, fmt.Sprintf("page not found: %s", r.URL.Path))
	}))
}

func (m *Module) WrapInMiddlewares(h http.Handler) http.Handler {
	return m.srv.WrapInMiddlewares(
		m.Middleware(
			h,
		),
	)
}
