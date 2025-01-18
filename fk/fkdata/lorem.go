package fkdata

import (
	"github.com/47monad/fakery/internal/binder"
	"golang.org/x/text/language"
)

func NewLorem(lang language.Tag) *binder.Data[Lorem] {
	d, err := binder.JSON[Lorem]("lorem", lang)
	if err != nil {
		panic(err)
	}
	return d
}

type Lorem struct {
	Words []string `json:"words"`
}

func KeyWord(d *Lorem) []string {
	return d.Words
}
