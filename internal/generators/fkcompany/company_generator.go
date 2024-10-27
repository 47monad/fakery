package fkcompany

import (
	"github.com/47monad/fakery/internal/fklocaler"
	"github.com/47monad/fakery/internal/fktemplater"
	"github.com/47monad/fakery/internal/fktypes"
	"github.com/47monad/fakery/internal/generators/fkperson"
)

type CompanyGenerator struct {
	Profession  func() string
	Suffix      func() string
	Industry    func() string
	Type        func() string
	Department  func() string
	BS          func() string
	Buzzword    func() string
	CatchPhrase func() string
	Name        func() string
}

func NewCompanyGenerator(conf *fktypes.FakeryConfig) CompanyGenerator {
	return CompanyGenerator{
		Profession:  generateProfession(conf),
		Suffix:      generateSuffix(conf),
		Industry:    generateIndustry(conf),
		Type:        generateType(conf),
		Department:  generateDepartment(conf),
		BS:          generateBS(conf),
		Buzzword:    generateBuzzword(conf),
		CatchPhrase: generateCatchPhrase(conf),
		Name:        generateName(conf),
	}
}

func generateProfession(conf *fktypes.FakeryConfig) func() string {
	return func() string {
		d := NewCompanyData(conf)
		_, sampler := fklocaler.Localize(d, keyProfession)
		return sampler()
	}
}

func generateDepartment(conf *fktypes.FakeryConfig) func() string {
	return func() string {
		d := NewCompanyData(conf)
		_, sampler := fklocaler.Localize(d, keyDepartment)
		return sampler()
	}
}

func generateSuffix(conf *fktypes.FakeryConfig) func() string {
	return func() string {
		d := NewCompanyData(conf)
		_, sampler := fklocaler.Localize(d, keySuffix)
		return sampler()
	}
}

func generateType(conf *fktypes.FakeryConfig) func() string {
	return func() string {
		d := NewCompanyData(conf)
		_, sampler := fklocaler.Localize(d, keyType)
		return sampler()
	}
}

func generateIndustry(conf *fktypes.FakeryConfig) func() string {
	return func() string {
		d := NewCompanyData(conf)
		_, sampler := fklocaler.Localize(d, keyIndustry)
		return sampler()
	}
}

func generateBuzzword(conf *fktypes.FakeryConfig) func() string {
	return func() string {
		d := NewCompanyData(conf)
		_, sampler := fklocaler.Localize(d, keyBuzzword)
		return sampler()
	}
}

func generateBS(conf *fktypes.FakeryConfig) func() string {
	return func() string {
		d := NewCompanyData(conf)
		_, sampler := fklocaler.Localize(d, keyBS)
		return sampler()
	}
}

func generateCatchPhrase(conf *fktypes.FakeryConfig) func() string {
	return func() string {
		d := NewCompanyData(conf)
		_, sampler := fklocaler.Localize(d, keyCatchPhrase)
		return sampler()
	}
}

// TODO. Fix locale usage in child generators. Currently we have hardcoded fallback language as below
func generateName(opts *fktypes.FakeryConfig) func() string {
	d := NewCompanyData(opts)
	pd := fkperson.NewPersonData(opts)
	return func() string {
		data, sampler := fklocaler.Localize(d, keyName)
		return fktemplater.MustExecTemplate("fullName", sampler(), generatorFuncMap(data, pd.Fallback))
	}
}
