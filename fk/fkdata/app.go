package fkdata

import (
	"github.com/47monad/fakery/internal/binder"
	"github.com/47monad/fakery/internal/locales"
	"golang.org/x/text/language"
)

type App struct {
	Name []string `json:"name"`
}

func NewApp(lang language.Tag) *binder.Data[App] {
	return locales.Load[App]("app", lang)
}
