package app

import (
	"html/template"
	"path/filepath"
)

const layout = "./web/*.layout.html"
const tmpl = "./web/*.tmpl.html"

var tmplFuncs = template.FuncMap{}

// Gen generates a template cache from the files that are inside ./web/
func Gen() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob(tmpl)
	if err != nil {
		return cache, err
	}

	for _, p := range pages {
		name := filepath.Base(p)
		tmplDef, err := template.New(name).Funcs(tmplFuncs).ParseFiles(p)
		if err != nil {
			return cache, err
		}

		match, err := filepath.Glob(layout)
		if err != nil {
			return cache, err
		}

		if len(match) > 0 {
			tmplDef, err = tmplDef.ParseGlob(layout)
			if err != nil {
				return cache, err
			}
		}
		cache[name] = tmplDef
	}
	return cache, err
}
