package fkinternet

import (
	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/fk/fkopts"
)

type Faker struct {
	ctx *fk.Context[fkdata.Internet]
}

func New(opts ...fkopts.Builder[fkopts.FakerOpts[fkdata.Internet]]) *Faker {
	_opts, _ := fkopts.Build(opts...)
	ctx := fk.NewContext(_opts.LW.GetFirstLang(), _opts.Data)
	return &Faker{ctx: ctx}
}
