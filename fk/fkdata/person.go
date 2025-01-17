package fkdata

import (
	"github.com/47monad/fakery/internal/binder"
	"golang.org/x/text/language"
)

type Person struct {
	MaleFirstName   []string `json:"maleFirstName"`
	FemaleFirstName []string `json:"femaleFirstName"`
	FirstName       []string `json:"firstName"`
	LastName        []string `json:"lastName"`
	NameWithMiddle  []string `json:"nameWithMiddle"`
	Prefix          []string `json:"prefix"`
	Suffix          []string `json:"suffix"`
	FullName        []string `json:"fullName"`
}

func NewPerson(lang language.Tag) *binder.Data[Person] {
	d, err := binder.JSON[Person]("person", lang)
	if err != nil {
		panic(err)
	}
	return d
}
