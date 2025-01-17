package fkcompany

import (
	"slices"

	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/fk/fkperson"
	"github.com/47monad/fakery/internal/sampler"
	"github.com/47monad/fakery/internal/templater"
)

func profession(ctx *fk.Context[fkdata.Company]) string {
	return sampler.Run(ctx.Data, func(d *fkdata.Company) []string { return d.Profession })
}

func department(ctx *fk.Context[fkdata.Company]) string {
	return sampler.Run(ctx.Data, func(d *fkdata.Company) []string { return d.Department })
}

func suffix(ctx *fk.Context[fkdata.Company]) string {
	return sampler.Run(ctx.Data, func(d *fkdata.Company) []string { return d.Suffix })
}

func cType(ctx *fk.Context[fkdata.Company]) string {
	return sampler.Run(ctx.Data, func(d *fkdata.Company) []string { return d.Type })
}

func industry(ctx *fk.Context[fkdata.Company]) string {
	return sampler.Run(ctx.Data, func(d *fkdata.Company) []string { return d.Industry })
}

func buzzword(ctx *fk.Context[fkdata.Company]) string {
	return sampler.Run(ctx.Data, func(d *fkdata.Company) []string { return slices.Concat(d.Buzzwords...) })
}

func bs(ctx *fk.Context[fkdata.Company]) string {
	return sampler.Run(ctx.Data, func(d *fkdata.Company) []string { return sampler.KeyMatrix(d.BS) })
}

func catchPhrase(ctx *fk.Context[fkdata.Company]) string {
	return sampler.Run(ctx.Data, func(d *fkdata.Company) []string { return sampler.KeyMatrix(d.Buzzwords) })
}

// TODO. Fix locale usage in child generators. Currently we have hardcoded fallback language as below
func name(ctx *fk.Context[fkdata.Company]) string {
	tmp := sampler.Run(ctx.Data, func(d *fkdata.Company) []string { return d.Name })
	personCtx := fk.NewContext[fkdata.Person](ctx.Lang, nil).SetBinder(fkdata.NewPerson)
	return templater.MustExec("companyName", tmp,
		templater.MergeFuncMaps(TemplateFuncMap(ctx), fkperson.TemplateFuncMap(personCtx, "male")),
	)
}
