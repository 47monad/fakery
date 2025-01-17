package fkphone

import (
	"html/template"

	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/fk/fkdata"
)

func TemplateFuncMap(ctx *fk.Context[fkdata.PhoneNumber]) template.FuncMap {
	return template.FuncMap{
		"areaCode":     func() string { return areaCode(ctx) },
		"exchangeCode": func() string { return exchangeCode(ctx) },
	}
}
