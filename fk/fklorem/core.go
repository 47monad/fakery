package fklorem

import (
	"strings"

	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/internal/sampler"
)

func word(ctx *fk.Context[fkdata.Lorem], exclude []string) string {
	return sampler.Run(ctx.Data, fkdata.KeyWord)
}

func words(ctx *fk.Context[fkdata.Lorem], count int) []string {
	if count == 0 || count == 1 {
		count = 3
	}
	return sampler.RunMany(count, ctx.Data, func(d *fkdata.Lorem) []string { return d.Words })
}

func sentence(ctx *fk.Context[fkdata.Lorem], wordCount int) string {
	if wordCount == 0 {
		wordCount = 6
	}
	return strings.Join(words(ctx, wordCount), " ")
}
