package fkperson

import (
	"github.com/47monad/fakery/internal/fklocaler"
	"github.com/47monad/fakery/internal/fktemplater"
	"github.com/47monad/fakery/internal/fktypes"
)

type PersonGenerator struct {
	FirstName      func(string) string
	LastName       func() string
	FullName       func(string) string
	NameWithMiddle func(string) string
	Prefix         func() string
	Suffix         func() string
}

type Person struct {
	FirstName  string
	MiddleName string
	LastName   string
	Prefix     string
	Suffix     string
}

func NewPersonGenerator(opts *fktypes.FakeryConfig) PersonGenerator {
	return PersonGenerator{
		FirstName:      generatePersonFirstName(opts),
		LastName:       generatePersonLastName(opts),
		FullName:       generatePersonName(opts),
		NameWithMiddle: generatePersonNameWithMiddle(opts),
		Prefix:         generatePersonPrefix(opts),
		Suffix:         generatePersonSuffix(opts),
	}
}

func generatePersonLastName(opts *fktypes.FakeryConfig) func() string {
	d := NewPersonData(opts)
	return func() string {
		_, sampler := fklocaler.Localize(d, keyLastName)
		return sampler()
	}
}

func generatePersonFirstName(opts *fktypes.FakeryConfig) func(string) string {
	d := NewPersonData(opts)
	return func(gender string) string {
		_, sampler := fklocaler.Localize(d, keyFirstName(gender))
		return sampler()
	}
}

func generatePersonName(opts *fktypes.FakeryConfig) func(string) string {
	d := NewPersonData(opts)
	return func(gender string) string {
		data, sampler := fklocaler.Localize(d, keyFullName)
		return fktemplater.MustExecTemplate("fullName", sampler(), generatorFuncMap(data, gender))
	}
}

func generatePersonNameWithMiddle(opts *fktypes.FakeryConfig) func(string) string {
	d := NewPersonData(opts)
	return func(gender string) string {
		data, sampler := fklocaler.Localize(d, keyNameWithMiddle)
		return fktemplater.MustExecTemplate("fullName", sampler(), generatorFuncMap(data, gender))
	}
}

func generatePersonPrefix(opts *fktypes.FakeryConfig) func() string {
	d := NewPersonData(opts)
	return func() string {
		_, sampler := fklocaler.Localize(d, keyPrefix)
		return sampler()
	}
}

func generatePersonSuffix(opts *fktypes.FakeryConfig) func() string {
	d := NewPersonData(opts)
	return func() string {
		_, sampler := fklocaler.Localize(d, keySuffix)
		return sampler()
	}
}
