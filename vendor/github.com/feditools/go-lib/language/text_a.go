package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextAccount returns a translated phrase.
func (l *Localizer) TextAccount(count int) *LocalizedString {
	lg := logger.WithField("func", "TextAccount")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Account",
			One:   "Account",
			Other: "Accounts",
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

// TextAddBlock returns a translated phrase.
func (l *Localizer) TextAddBlock(count int) *LocalizedString {
	lg := logger.WithField("func", "TextAddBlock")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "AddBlock",
			One:   "Add Block",
			Other: "Add Blocks",
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

// TextAddOauth20Client returns a translated phrase.
func (l *Localizer) TextAddOauth20Client(count int) *LocalizedString {
	lg := logger.WithField("func", "TextAddOauth20Client")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "AddOauth20Client",
			One:   "Add OAuth 2.0 Client",
			Other: "Add OAuth 2.0 Clients",
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

// TextAdmin returns a translated phrase.
func (l *Localizer) TextAdmin() *LocalizedString {
	lg := logger.WithField("func", "TextAdmin")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Admin",
			Other: "Admin",
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

// TextAllow returns a translated phrase.
func (l *Localizer) TextAllow() *LocalizedString {
	lg := logger.WithField("func", "TextAllow")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Allow",
			Other: "Allow",
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

// TextApplicationToken returns a translated phrase.
func (l *Localizer) TextApplicationToken(count int) *LocalizedString {
	lg := logger.WithField("func", "TextApplicationToken")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "ApplicationToken",
			One:   "Application Token",
			Other: "Application Tokens",
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

// TextAuthorize returns a translated phrase.
func (l *Localizer) TextAuthorize() *LocalizedString {
	lg := logger.WithField("func", "TextAuthorize")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Authorize",
			Other: "Authorize",
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

// TextAuthorizeApplicationDescription returns a translated phrase.
func (l *Localizer) TextAuthorizeApplicationDescription(description string) *LocalizedString {
	lg := logger.WithField("func", "TextAuthorizeApplicationDescription")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "AuthorizeApplicationDescription",
			Other: "Authorize {{.Description}}",
		},
		TemplateData: map[string]interface{}{
			"Description": description,
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
