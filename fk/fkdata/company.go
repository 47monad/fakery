package fkdata

import (
	"github.com/47monad/fakery/internal/binder"
	"github.com/47monad/fakery/internal/locales"
	"golang.org/x/text/language"
)

type Company struct {
	Suffix     []string   `json:"suffix"`
	Buzzwords  [][]string `json:"buzzwords"`
	BS         [][]string `json:"bs"`
	Name       []string   `json:"name"`
	Profession []string   `json:"profession"`
	Type       []string   `json:"type"`
	Industry   []string   `json:"industry"`
	Department []string   `json:"department"`
}

func NewCompany(lang language.Tag) *binder.Data[Company] {
	return locales.Load[Company]("company", lang)
}
