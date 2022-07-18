package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextSystem returns a translated phrase.
func (l *Localizer) TextSystem(count int) *LocalizedString {
	lg := logger.WithField("func", "TextSystem")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "System",
			One:   "System",
			Other: "Systems",
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
