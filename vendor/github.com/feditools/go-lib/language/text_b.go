package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextBlock returns a translated phrase.
func (l *Localizer) TextBlock(count int) *LocalizedString {
	lg := logger.WithField("func", "TextBlock")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Block",
			One:   "Block",
			Other: "Blocks",
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

// TextBlockedInstance returns a translated phrase.
func (l *Localizer) TextBlockedInstance(count int) *LocalizedString {
	lg := logger.WithField("func", "TextBlockedInstance")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "BlockedInstance",
			One:   "Blocked Instance",
			Other: "Blocked Instances",
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
