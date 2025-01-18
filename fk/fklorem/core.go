package fklorem

import (
	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/internal/sampler"
)

func word(ctx *fk.Context[fkdata.Lorem], exclude []string) string {
	return sampler.Run(ctx.Data, fkdata.KeyWord)
}
