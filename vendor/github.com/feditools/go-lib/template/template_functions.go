package template

import "html/template"

const (
	funcNameDec      = "dec"
	funcNameHTMLSafe = "htmlSafe"
	funcNameJSSafe   = "jsSafe"
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
	funcJSSafe = func(javascript string) template.JS {
		/* #nosec G203 */
		return template.JS(javascript)
	}
	funcInc = func(i int) int {
		i++

		return i
	}

	defaultFunctions = template.FuncMap{
		funcNameDec:      funcDec,
		funcNameHTMLSafe: funcHTMLSafe,
		funcNameJSSafe:   funcJSSafe,
		funcNameInc:      funcInc,
	}
)
