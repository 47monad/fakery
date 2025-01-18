package fkperson

import (
	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/fk/fkopts"
)

type Faker struct {
	ctx *fk.Context[fkdata.Person]
}

func New(opts ...fkopts.Builder[fkopts.FakerOpts[fkdata.Person]]) *Faker {
	_opts, _ := fkopts.Build(opts...)
	ctx := fk.NewContext(_opts.LW.GetFirstLang(), _opts.Data)
	return &Faker{ctx: ctx}
}
