package template

import "html/template"

const (
	funcNameDec      = "dec"
	funcNameHTMLSafe = "htmlSafe"
	funcNameInc      = "inc"
)

var (
	funcDec = func(i int) int {
		i--

		return i
	}
	funcHTMLSafe = func(html string) template.HTML {
		/* #nosec G203 */
		return template.HTML(html)
	}
	funcInc = func(i int) int {
		i++

		return i
	}

	defaultFunctions = template.FuncMap{
		funcNameDec:      funcDec,
		funcNameHTMLSafe: funcHTMLSafe,
		funcNameInc:      funcInc,
	}
)
