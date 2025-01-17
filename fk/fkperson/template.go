package fkperson

import (
	"text/template"

	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/fk/fkdata"
)

func TemplateFuncMap(ctx *fk.Context[fkdata.Person], gender string) template.FuncMap {
	return template.FuncMap{
		"prefix":    func() string { return prefix(ctx) },
		"suffix":    func() string { return suffix(ctx) },
		"firstName": func() string { return firstName(ctx, gender) },
		"lastName":  func() string { return lastName(ctx) },
	}
}
