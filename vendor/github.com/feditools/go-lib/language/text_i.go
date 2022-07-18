package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextInstance returns a translated phrase.
func (l *Localizer) TextInstance(count int) *LocalizedString {
	lg := logger.WithField("func", "TextInstance")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Instance",
			One:   "Instance",
			Other: "Instances",
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

// TextInvalidURI returns a translated phrase.
func (l *Localizer) TextInvalidURI(count int) *LocalizedString {
	lg := logger.WithField("func", "TextInvalidURI")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "InvalidURI",
			One:   "Invalid URI",
			Other: "Invalid URIs",
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
