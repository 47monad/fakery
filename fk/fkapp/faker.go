package fkapp

import (
	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/fk/fkopts"
)

type Faker struct {
	ctx *fk.Context[fkdata.App]
}

func New(opts ...fkopts.Builder[fkopts.FakerOpts[fkdata.App]]) *Faker {
	_opts, _ := fkopts.Build(opts...)
	ctx := fk.NewContext(_opts.LW.GetFirstLang(), _opts.Data).SetBinder(fkdata.NewApp)
	return &Faker{ctx: ctx}
}

func (f *Faker) SemVer(opts ...fkopts.Builder[fkopts.SemVerOpts]) string {
	_opts, _ := fkopts.Build(opts...)
	return semanticVersion(_opts)
}

func (f *Faker) Name(opts ...fkopts.Builder[fkopts.AppNameOpts]) string {
	//_opts, _ := fkopts.Build(opts...)
	// ctx := fk.NewContext(_opts.LW.GetFirstLang(), _opts.Data).SetBinder(fkdata.NewApp)
	return name(f.ctx)
}
