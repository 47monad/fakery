package fkdata

import (
	"github.com/47monad/fakery/internal/binder"
	"github.com/47monad/fakery/internal/locales"
	"golang.org/x/text/language"
)

func NewLorem(lang language.Tag) *binder.Data[Lorem] {
	return locales.Load[Lorem]("lorem", lang)
}

type Lorem struct {
	Words []string `json:"words"`
}

func KeyWord(d *Lorem) []string {
	return d.Words
}
