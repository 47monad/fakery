package fkinternet

import (
	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/fk/fkdata"
	"golang.org/x/text/language"
)

func Domain() string {
	ctx := fk.NewContext[fkdata.Internet](language.English, nil).SetBinder(fkdata.NewInternet)
	return domainName(ctx)
}
