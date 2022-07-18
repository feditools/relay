package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextOauth returns a translated phrase.
func (l *Localizer) TextOauth() *LocalizedString {
	lg := logger.WithField("func", "TextOauth")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Oauth",
			Other: "OAuth",
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

// TextOauth20Client returns a translated phrase.
func (l *Localizer) TextOauth20Client(count int) *LocalizedString {
	lg := logger.WithField("func", "TextOauth20Client")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Oauth20Client",
			One:   "OAuth 2.0 Client",
			Other: "OAuth 2.0 Clients",
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

// TextOauth20Settings returns a translated phrase.
func (l *Localizer) TextOauth20Settings() *LocalizedString {
	lg := logger.WithField("func", "TextOauth20Settings")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Oauth20Settings",
			Other: "OAuth 2.0 Settings",
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
