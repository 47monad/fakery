package fkphone

import (
	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/fk/fkopts"
)

func Number(opts ...fkopts.Builder[fkopts.GeneralOpts[fkdata.PhoneNumber]]) string {
	_opts, _ := fkopts.Build(opts...)
	ctx := fk.NewContext(_opts.LW.GetFirstLang(), _opts.Data).SetBinder(fkdata.NewPhoneNumber)
	return number(ctx)
}

func MobileNumber(opts ...fkopts.Builder[fkopts.GeneralOpts[fkdata.PhoneNumber]]) string {
	_opts, _ := fkopts.Build(opts...)
	ctx := fk.NewContext(_opts.LW.GetFirstLang(), _opts.Data).SetBinder(fkdata.NewPhoneNumber)
	return mobileNumber(ctx)
}

func E164(opts ...fkopts.Builder[fkopts.GeneralOpts[fkdata.PhoneNumber]]) string {
	_opts, _ := fkopts.Build(opts...)
	ctx := fk.NewContext(_opts.LW.GetFirstLang(), _opts.Data).SetBinder(fkdata.NewPhoneNumber)
	return e164(ctx)
}
