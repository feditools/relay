package fedihelper

import "strings"

type WebfingerURI string

func (w WebfingerURI) FTemplate() string {
	return strings.ReplaceAll(string(w), "{uri}", "acct:%s@%s")
}

func (w WebfingerURI) String() string {
	return string(w)
}
