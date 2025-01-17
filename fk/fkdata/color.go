package fkdata

import (
	"github.com/47monad/fakery/internal/binder"
	"golang.org/x/text/language"
)

type Color struct {
	Name []string `json:"name"`
}

func NewColor(lang language.Tag) *binder.Data[Color] {
	d, err := binder.JSON[Color]("color", lang)
	if err != nil {
		panic(err)
	}
	return d
}
