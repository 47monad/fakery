package fkdata

import (
	"github.com/47monad/fakery/internal/binder"
	"golang.org/x/text/language"
)

type App struct {
	Name []string `json:"name"`
}

func NewApp(lang language.Tag) *binder.Data[App] {
	d, err := binder.JSON[App]("app", lang)
	if err != nil {
		panic(err)
	}
	return d
}
