package fkdata

import (
	"github.com/47monad/fakery/internal/binder"
	"github.com/47monad/fakery/internal/locales"
	"golang.org/x/text/language"
)

type Color struct {
	Name []string `json:"name"`
}

func NewColor(lang language.Tag) *binder.Data[Color] {
	return locales.Load[Color]("color", lang)
}
