package fkdevice

import (
	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/fk/fkopts"
	"github.com/47monad/fakery/internal/sampler"
)

type Faker struct {
	ctx *fk.Context[fkdata.Device]
}

func New(opts ...fkopts.Builder[fkopts.FakerOpts[fkdata.Device]]) *Faker {
	_opts, err := fkopts.Build(opts...)
	if err != nil {
		panic(err)
	}
	ctx := fk.NewContext(_opts.LW.GetFirstLang(), _opts.Data).SetBinder(fkdata.NewDevice)
	return &Faker{ctx: ctx}
}

func (f *Faker) Model() string {
	return model(f.ctx)
}

func (f *Faker) Manufacturer() string {
	return manufacturer(f.ctx)
}

func (f *Faker) Platform() string {
	return platform(f.ctx)
}

func (f *Faker) Type() string {
	return dType(f.ctx)
}

func (f *Faker) SerialNumber() string {
	return serialNumber(f.ctx)
}

func model(ctx *fk.Context[fkdata.Device]) string {
	return sampler.Run(ctx.Data, func(d *fkdata.Device) []string { return d.ModelName })
}

func manufacturer(ctx *fk.Context[fkdata.Device]) string {
	return sampler.Run(ctx.Data, func(d *fkdata.Device) []string { return d.Manufacturer })
}

func platform(ctx *fk.Context[fkdata.Device]) string {
	return sampler.Run(ctx.Data, func(d *fkdata.Device) []string { return d.Platform })
}

func dType(ctx *fk.Context[fkdata.Device]) string {
	return sampler.Run(ctx.Data, func(d *fkdata.Device) []string { return d.Type })
}

func serialNumber(ctx *fk.Context[fkdata.Device]) string {
	return sampler.Run(ctx.Data, func(d *fkdata.Device) []string { return d.SerialNumber })
}
