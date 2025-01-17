package fkperson

import (
	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/internal/sampler"
	"github.com/47monad/fakery/internal/templater"
)

func lastName(ctx *fk.Context[fkdata.Person]) string {
	return sampler.Run(ctx.Data, func(d *fkdata.Person) []string { return d.LastName })
}

func firstName(ctx *fk.Context[fkdata.Person], gender string) string {
	keyer := func(d *fkdata.Person) []string {
		if gender == "female" {
			return d.FemaleFirstName
		}
		return d.MaleFirstName
	}
	return sampler.Run(ctx.Data, keyer)
}

func name(ctx *fk.Context[fkdata.Person], gender string) string {
	tmpl := sampler.Run(ctx.Data, func(d *fkdata.Person) []string { return d.FullName })
	return templater.MustExec("fullName", tmpl, TemplateFuncMap(ctx, gender))
}

func nameWithMiddle(ctx *fk.Context[fkdata.Person], gender string) string {
	tmpl := sampler.Run(ctx.Data, func(d *fkdata.Person) []string { return d.NameWithMiddle })
	return templater.MustExec("fullName", tmpl, TemplateFuncMap(ctx, gender))
}

func prefix(ctx *fk.Context[fkdata.Person]) string {
	return sampler.Run(ctx.Data, func(d *fkdata.Person) []string { return d.Prefix })
}

func suffix(ctx *fk.Context[fkdata.Person]) string {
	return sampler.Run(ctx.Data, func(d *fkdata.Person) []string { return d.Suffix })
}
