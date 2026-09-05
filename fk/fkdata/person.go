package fkdata

import (
	"github.com/47monad/fakery/internal/binder"
	"github.com/47monad/fakery/internal/locales"
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
	return locales.Load[Person]("person", lang)
}
