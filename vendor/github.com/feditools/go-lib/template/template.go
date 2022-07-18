package template

import (
	"embed"
	"html/template"
	"io/ioutil"
	"strings"
)

const templateDir = "templates"

// templateFS contains template files required by the application
//go:embed templates/*.gohtml
var templateFS embed.FS

// New creates a new template.
func New(f template.FuncMap) (*template.Template, error) {
	tpl := template.New("")

	tmplFuncs := template.FuncMap{}
	for k, v := range defaultFunctions {
		tmplFuncs[k] = v
	}
	if len(f) > 0 {
		for k, v := range f {
			tmplFuncs[k] = v
		}
	}
	tpl.Funcs(tmplFuncs)

	dir, err := templateFS.ReadDir(templateDir)
	if err != nil {
		return nil, err
	}
	for _, d := range dir {
		filePath := templateDir + "/" + d.Name()
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".gohtml") {
			continue
		}

		// open it
		file, err := templateFS.Open(filePath)
		if err != nil {
			return nil, err
		}

		// read it
		tmplData, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}

		// It can now be parsed as a string.
		_, err = tpl.Parse(string(tmplData))
		if err != nil {
			return nil, err
		}
	}

	return tpl, nil
}
