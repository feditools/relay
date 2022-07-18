package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextRelay returns a translated phrase.
func (l *Localizer) TextRelay(count int) *LocalizedString {
	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Relay",
			One:   "Relay",
			Other: "Relays",
		},
		PluralCount: count,
	})
	if err != nil {
		logger.WithField("func", "TextRelay").Warningf(missingTranslationWarning, err.Error())
	}

	return &LocalizedString{
		language: tag,
		string:   text,
	}
}

// TextRequired returns a translated phrase.
func (l *Localizer) TextRequired() *LocalizedString {
	lg := logger.WithField("func", "TextRequired")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Required",
			Other: "Required",
		},
	})
	if err != nil {
		lg.Warningf(missingTranslationWarning, err.Error())
	}

	return &LocalizedString{
		language: tag,
		string:   text,
	}
}

// TextRedirectURI returns a translated phrase.
func (l *Localizer) TextRedirectURI(count int) *LocalizedString {
	lg := logger.WithField("func", "TextRedirectURI")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "RedirectURI",
			One:   "Redirect URI",
			Other: "Redirect URIs",
		},
		PluralCount: count,
	})
	if err != nil {
		lg.Warningf(missingTranslationWarning, err.Error())
	}

	return &LocalizedString{
		language: tag,
		string:   text,
	}
}
