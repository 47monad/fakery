package fkcompany

import (
	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/fk/fkdata"
	"golang.org/x/text/language"
)

func Buzzword() string {
	ctx := fk.NewContext[fkdata.Company](language.English, nil).SetBinder(fkdata.NewCompany)
	return buzzword(ctx)
}
