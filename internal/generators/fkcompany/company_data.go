package fkcompany

import (
	"slices"

	"github.com/47monad/fakery/internal/fkgeneral"
	"github.com/47monad/fakery/internal/fkloader"
	"github.com/47monad/fakery/internal/fktypes"
)

type CompanyData struct {
	Suffix     []string   `json:"suffix"`
	Buzzwords  [][]string `json:"buzzwords"`
	BS         [][]string `json:"bs"`
	Name       []string   `json:"name"`
	Profession []string   `json:"profession"`
	Type       []string   `json:"type"`
	Industry   []string   `json:"industry"`
	Department []string   `json:"department"`
}

func NewCompanyData(conf *fktypes.FakeryConfig) *fktypes.ResultData[CompanyData] {
	d, err := fkloader.FromJSON[CompanyData](conf.FS, "company", conf.DefaultLang)
	if err != nil {
		panic(err)
	}
	return d
}

func keyProfession(d *CompanyData) []string {
	return d.Profession
}

func keyDepartment(d *CompanyData) []string {
	return d.Department
}

func keySuffix(d *CompanyData) []string {
	return d.Suffix
}

func keyType(d *CompanyData) []string {
	return d.Type
}

func keyIndustry(d *CompanyData) []string {
	return d.Industry
}

func keyBuzzword(d *CompanyData) []string {
	return slices.Concat(d.Buzzwords...)
}

func keyBS(d *CompanyData) []string {
	return fkgeneral.ResampleString(d.BS)
}

func keyCatchPhrase(d *CompanyData) []string {
	return fkgeneral.ResampleString(d.Buzzwords)
}

func keyName(d *CompanyData) []string {
	return d.Name
}
