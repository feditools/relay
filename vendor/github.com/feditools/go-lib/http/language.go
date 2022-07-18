package http

import (
	"golang.org/x/text/language"
)

func GetPageLang(query, header, defaultLang string) string {
	if query != "" {
		t, _, err := language.ParseAcceptLanguage(query)
		if err == nil {
			if len(t) > 0 {
				return t[0].String()
			}
		}
	}

	if header != "" {
		t, _, err := language.ParseAcceptLanguage(header)
		if err == nil {
			if len(t) > 0 {
				return t[0].String()
			}
		}
	}

	return defaultLang
}
