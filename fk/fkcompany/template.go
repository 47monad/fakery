package fkcompany

import (
	"html/template"

	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/fk/fkdata"
)

func TemplateFuncMap(ctx *fk.Context[fkdata.Company]) template.FuncMap {
	return template.FuncMap{
		"suffix": func() string { return suffix(ctx) },
	}
}
