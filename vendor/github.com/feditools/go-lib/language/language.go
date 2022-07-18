package language

import (
	"embed"
	"io/ioutil"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

// Locales contains static files required by the application
//go:embed locales/active.*.yaml
var Locales embed.FS

// DefaultLanguage is the default language of the application.
var DefaultLanguage = language.English

// Module represent the language module for translating text.
type Module struct {
	lang       language.Tag
	langBundle *i18n.Bundle
}

// New creates a new language module.
func New() (*Module, error) {
	l := logger.WithField("func", "New")

	module := Module{
		lang:       DefaultLanguage,
		langBundle: i18n.NewBundle(DefaultLanguage),
	}

	module.langBundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	dir, err := Locales.ReadDir("locales")
	if err != nil {
		return nil, err
	}
	for _, d := range dir {
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".yaml") {
			continue
		}
		l.Debugf("loading language file: %s", d.Name())

		// open it
		file, err := Locales.Open("locales/" + d.Name())
		if err != nil {
			return nil, err
		}

		// read it
		buffer, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}

		// paris if not empty
		empty := isEmptyYaml(buffer)
		if !empty {
			module.langBundle.MustParseMessageFileBytes(buffer, d.Name())
		}
	}

	return &module, nil
}

// Language returns the default language.
func (m Module) Language() language.Tag { return m.lang }
