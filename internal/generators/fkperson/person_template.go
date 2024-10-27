package fkperson

import (
	"text/template"

	"github.com/47monad/fakery/internal/fklocaler"
)

func generatorFuncMap(d *PersonData, gender string) template.FuncMap {
	return template.FuncMap{
		"prefix":    WithData(d, prefixGFunc),
		"suffix":    WithData(d, suffixGFunc),
		"firstName": WithData(d, firstNameGGFunc(gender)),
		"lastName":  WithData(d, LastNameGFunc),
	}
}

func WithData(d *PersonData, f func(*PersonData, ...interface{}) (interface{}, error)) func(...interface{}) (interface{}, error) {
	return func(args ...interface{}) (interface{}, error) {
		return f(d, args)
	}
}
func firstNameGGFunc(gender string) func(*PersonData, ...interface{}) (interface{}, error) {
	return func(d *PersonData, args ...interface{}) (interface{}, error) {
		return fklocaler.Sample(d, keyFirstName(gender)), nil
	}
}
func LastNameGFunc(d *PersonData, args ...interface{}) (interface{}, error) {
	return fklocaler.Sample(d, keyLastName), nil
}
func prefixGFunc(d *PersonData, args ...interface{}) (interface{}, error) {
	return fklocaler.Sample(d, keyPrefix), nil
}
func suffixGFunc(d *PersonData, args ...interface{}) (interface{}, error) {
	return fklocaler.Sample(d, keySuffix), nil
}
