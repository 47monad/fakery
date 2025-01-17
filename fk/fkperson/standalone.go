package fkperson

import (
	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/fk/fkopts"
)

func FirstName(opts ...fkopts.Builder[fkopts.NameOpts]) string {
	_opts, _ := fkopts.Build(opts...)
	ctx := fk.NewContext(_opts.Tag, _opts.Data).SetBinder(fkdata.NewPerson)
	return firstName(ctx, _opts.Gender)
}

func LastName(opts ...fkopts.Builder[fkopts.LastNameOpts]) string {
	_opts, _ := fkopts.Build(opts...)
	ctx := fk.NewContext(_opts.Tag, _opts.Data).SetBinder(fkdata.NewPerson)
	return lastName(ctx)
}
