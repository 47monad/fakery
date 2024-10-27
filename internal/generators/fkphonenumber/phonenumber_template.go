package fkphonenumber

import (
	"text/template"

	"github.com/47monad/fakery/internal/fklocaler"
	"github.com/47monad/fakery/internal/generators/fknumber"
)

func PhoneNumberGFunc(d *PhoneNumberData) template.FuncMap {
	return template.FuncMap{
		"num":          fknumber.NumTFunc,
		"areaCode":     WithPhoneNumberData(d, areaCodeGFunc),
		"exchangeCode": WithPhoneNumberData(d, exchangeCodeGFunc),
	}
}

func WithPhoneNumberData(d *PhoneNumberData, f func(*PhoneNumberData, ...interface{}) (interface{}, error)) func(...interface{}) (interface{}, error) {
	return func(args ...interface{}) (interface{}, error) {
		return f(d, args)
	}
}

func areaCodeGFunc(d *PhoneNumberData, args ...interface{}) (interface{}, error) {
	return fklocaler.Sample(d, keyPhoneNumberAreaCode), nil
}

func exchangeCodeGFunc(d *PhoneNumberData, args ...interface{}) (interface{}, error) {
	return fklocaler.Sample(d, keyPhoneNumberExchangeCode), nil
}
