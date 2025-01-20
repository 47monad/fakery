package fkapp

import (
	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/fk/fkopts"
)

func SemVer(opts ...fkopts.Builder[fkopts.SemVerOpts]) string {
	_opts, _ := fkopts.Build(opts...)
	return semanticVersion(_opts)
}

func Name(opts ...fkopts.Builder[fkopts.AppNameOpts]) string {
	_opts, _ := fkopts.Build(opts...)
	ctx := fk.NewContext(_opts.LW.GetFirstLang(), _opts.Data).SetBinder(fkdata.NewApp)
	return name(ctx)
}
