package fkapp

import (
	"fmt"

	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/fk/fkopts"
	"github.com/47monad/fakery/internal/randomizer"
	"github.com/47monad/fakery/internal/sampler"
)

func semanticVersion(opts *fkopts.SemVerOpts) string {
	return fmt.Sprintf("%d.%d.%d",
		randomizer.InRange(opts.Major.Min, opts.Major.Max),
		randomizer.InRange(opts.Minor.Min, opts.Minor.Max),
		randomizer.InRange(opts.Patch.Min, opts.Patch.Max))
}

func name(ctx *fk.Context[fkdata.App]) string {
	return sampler.Run(ctx.Data, func(d *fkdata.App) []string { return d.Name })
}
