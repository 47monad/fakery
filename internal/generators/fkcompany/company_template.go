package fkcompany

import (
	"text/template"

	"github.com/47monad/fakery/internal/fklocaler"
	"github.com/47monad/fakery/internal/generators/fkperson"
)

func generatorFuncMap(d *CompanyData, pd *fkperson.PersonData) template.FuncMap {
	return template.FuncMap{
		"suffix":   WithData(d, SuffixGFunc),
		"lastName": fkperson.WithData(pd, fkperson.LastNameGFunc),
	}
}

func WithData(d *CompanyData, f func(*CompanyData, ...interface{}) (interface{}, error)) func(...interface{}) (interface{}, error) {
	return func(args ...interface{}) (interface{}, error) {
		return f(d, args)
	}
}

func SuffixGFunc(d *CompanyData, args ...interface{}) (interface{}, error) {
	return fklocaler.Sample(d, keySuffix), nil
}
