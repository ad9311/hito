package handler

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	"github.com/ad9311/hito/internal/app"
	"github.com/ad9311/hito/internal/console"
)

func writeTemplate(w http.ResponseWriter, tmpl string) {
	templateMap := defaultCache()
	t, exist := templateMap[tmpl]
	if !exist {
		console.AssertFatal(fmt.Errorf("template %s does not exist", tmpl))
	}

	buff := new(bytes.Buffer)
	err := t.Execute(buff, data)
	console.AssertPanic(err)

	_, err = buff.WriteTo(w)
	console.AssertPanic(err)
}

func defaultCache() map[string]*template.Template {
	if config.UseCache {
		return config.TemplateCache
	}
	templateCache, err := app.Gen()
	console.AssertError(err)
	config.TemplateCache = templateCache
	return templateCache
}
