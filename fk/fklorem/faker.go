package fklorem

import (
	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/fk/fkopts"
	"github.com/47monad/fakery/internal/sampler"
)

type Faker struct {
	ctx *fk.Context[fkdata.Lorem]
}

func New(opts ...fkopts.Builder[fkopts.FakerOpts]) *Faker {
	_opts, _ := fkopts.Build(opts...)
	ctx := fk.NewContext(_opts.Tag, _opts.Data)
	if ctx.Data == nil {
		ctx.Data = fkdata.NewBinder(ctx.Lang)
	}
	return &Faker{ctx: ctx}
}

func (f *Faker) Word(opts ...fkopts.Builder[fkopts.WordOpts]) string {
	_opts, _ := fkopts.Build(opts...)
	return word(f.ctx, _opts.ExcludedWords)
}

func Word(opts ...fkopts.Builder[fkopts.WordOpts]) string {
	_opts, _ := fkopts.Build(opts...)
	ctx := fk.NewContext(_opts.Tag, _opts.Data)
	return word(ctx, _opts.ExcludedWords)
}

func word(ctx *fk.Context[fkdata.Lorem], exclude []string) string {
	data := ctx.Data
	if ctx.Data == nil {
		data = fkdata.NewBinder(ctx.Lang)
	}
	return sampler.Run(data, fkdata.KeyWord)
}
