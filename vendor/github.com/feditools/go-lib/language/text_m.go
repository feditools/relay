package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextModeration returns a translated phrase.
func (l *Localizer) TextModeration() *LocalizedString {
	lg := logger.WithField("func", "TextLogin")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Moderation",
			Other: "Moderation",
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
