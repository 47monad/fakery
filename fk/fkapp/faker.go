package fkapp

import (
	"fmt"

	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/fk/fkopts"
	"github.com/47monad/fakery/internal/randomizer"
	"github.com/47monad/fakery/internal/sampler"
)

type Faker struct {
	ctx *fk.Context[fkdata.App]
}

func New(opts ...fkopts.Builder[fkopts.FakerOpts[fkdata.App]]) *Faker {
	_opts, _ := fkopts.Build(opts...)
	ctx := fk.NewContext(_opts.LW.GetFirstLang(), _opts.Data).SetBinder(fkdata.NewApp)
	return &Faker{ctx: ctx}
}

func (f *Faker) SemanticVersion(opts ...fkopts.Builder[fkopts.SemVerOpts]) string {
	_opts, _ := fkopts.Build(opts...)
	return semanticVersion(_opts)
}

func (f *Faker) Name() string {
	return name(f.ctx)
}

func semanticVersion(opts *fkopts.SemVerOpts) string {
	return fmt.Sprintf("%d.%d.%d",
		randomizer.InRange(opts.Major.Min, opts.Major.Max),
		randomizer.InRange(opts.Minor.Min, opts.Minor.Max),
		randomizer.InRange(opts.Patch.Min, opts.Patch.Max))
}

func name(ctx *fk.Context[fkdata.App]) string {
	return sampler.Run(ctx.Data, func(d *fkdata.App) []string { return d.Name })
}
