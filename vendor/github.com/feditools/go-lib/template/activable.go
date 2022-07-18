package template

import (
	"regexp"
)

// ActivableSlice a slice where each element has an active bit which can be set based on a string.
type ActivableSlice interface {
	GetChildren(i int) ActivableSlice
	GetMatcher(i int) *regexp.Regexp
	SetActive(i int, a bool)
	Len() int
}

// SetActive sets an active bit in a slice.
func SetActive(a ActivableSlice, s string) {
	for i := 0; i < a.Len(); i++ {
		matcher := a.GetMatcher(i)
		if matcher != nil {
			matcherFound := matcher.Match([]byte(s))
			if matcherFound {
				a.SetActive(i, true)
			}
		}
		children := a.GetChildren(i)
		if children != nil {
			SetActive(children, s)
		}
	}
}
