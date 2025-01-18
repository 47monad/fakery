package fklorem

import (
	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/fk/fkopts"
)

type Faker struct {
	ctx *fk.Context[fkdata.Lorem]
}

func New(opts ...fkopts.Builder[fkopts.FakerOpts[fkdata.Lorem]]) *Faker {
	_opts, _ := fkopts.Build(opts...)
	ctx := fk.NewContext(_opts.LW.GetFirstLang(), _opts.Data).SetBinder(fkdata.NewLorem)
	return &Faker{ctx: ctx}
}

func (f *Faker) Word(opts ...fkopts.Builder[fkopts.WordOpts]) string {
	_opts, _ := fkopts.Build(opts...)
	return word(f.ctx, _opts.ExcludedWords)
}
