package fkperson

import (
	"github.com/47monad/fakery/internal/fkloader"
	"github.com/47monad/fakery/internal/fktypes"
)

type PersonData struct {
	MaleFirstName   []string `json:"maleFirstName"`
	FemaleFirstName []string `json:"femaleFirstName"`
	FirstName       []string `json:"firstName"`
	LastName        []string `json:"lastName"`
	NameWithMiddle  []string `json:"nameWithMiddle"`
	Prefix          []string `json:"prefix"`
	Suffix          []string `json:"suffix"`
	FullName        []string `json:"fullName"`
}

func NewPersonData(opts *fktypes.FakeryConfig) *fktypes.ResultData[PersonData] {
	d, err := fkloader.FromJSON[PersonData](opts.FS, "person", opts.DefaultLang)
	if err != nil {
		panic(err)
	}
	return d
}

func keyFirstName(gender string) func(*PersonData) []string {
	return func(p *PersonData) []string {
		if gender == "female" {
			return keyFemaleFirstName(p)
		}
		return keyMaleFirstName(p)
	}
}

func keyMaleFirstName(d *PersonData) []string {
	return d.MaleFirstName
}

func keyFemaleFirstName(d *PersonData) []string {
	return d.MaleFirstName
}

func keyLastName(d *PersonData) []string {
	return d.LastName
}

func keyFullName(d *PersonData) []string {
	return d.FullName
}

func keyNameWithMiddle(d *PersonData) []string {
	return d.NameWithMiddle
}

func keyPrefix(d *PersonData) []string {
	return d.Prefix
}

func keySuffix(d *PersonData) []string {
	return d.Suffix
}
