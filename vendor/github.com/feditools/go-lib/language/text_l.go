package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextList returns a translated phrase.
func (l *Localizer) TextList(count int) *LocalizedString {
	lg := logger.WithField("func", "TextList")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "List",
			One:   "List",
			Other: "Lists",
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

// TextLogin returns a translated phrase.
func (l *Localizer) TextLogin() *LocalizedString {
	lg := logger.WithField("func", "TextLogin")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Login",
			Other: "Login",
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

// TextLogout returns a translated phrase.
func (l *Localizer) TextLogout() *LocalizedString {
	lg := logger.WithField("func", "TextLogout")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Logout",
			Other: "Logout",
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

// TextLooksGood returns a translated phrase.
func (l *Localizer) TextLooksGood() *LocalizedString {
	lg := logger.WithField("func", "TextLooksGood")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "LooksGood",
			Other: "Looks Good!",
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
