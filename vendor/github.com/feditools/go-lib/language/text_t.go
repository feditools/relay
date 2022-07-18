package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextToken returns a translated phrase.
func (l *Localizer) TextToken(count int) *LocalizedString {
	lg := logger.WithField("func", "TextSystem")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Token",
			One:   "Token",
			Other: "Tokens",
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
