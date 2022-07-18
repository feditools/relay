package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextFediverse returns a translated phrase.
func (l *Localizer) TextFediverse() *LocalizedString {
	lg := logger.WithField("func", "TextFediverse")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Fediverse",
			Other: "Fediverse",
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
