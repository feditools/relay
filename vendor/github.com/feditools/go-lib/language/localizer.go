package language

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// Localizer returns translated phrases.
type Localizer struct {
	localizer *i18n.Localizer
}

// NewLocalizer returns a localizer which will return translated phrases based on the provided languages.
func (m Module) NewLocalizer(langs ...string) (*Localizer, error) {
	return &Localizer{
		localizer: i18n.NewLocalizer(m.langBundle, langs...),
	}, nil
}

// LocalizedString contains a localized string.
type LocalizedString struct {
	language language.Tag
	string   string
}

// Language returns the language of the localized string.
func (l *LocalizedString) Language() language.Tag { return l.language }

// String returns the localized string.
func (l *LocalizedString) String() string { return l.string }
