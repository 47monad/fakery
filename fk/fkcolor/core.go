package fkcolor

import (
	"fmt"

	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/internal/randomizer"
	"github.com/47monad/fakery/internal/sampler"
)

func name(ctx *fk.Context[fkdata.Color]) string {
	return sampler.Run(ctx.Data, func(d *fkdata.Color) []string { return d.Name })
}

func rgbColor() [3]int {
	num := func() int { return randomizer.InRange(0, 255) }

	return [3]int{num(), num(), num()}
}

func hexColor() string {
	rgb := rgbColor()

	return fmt.Sprintf("#%x%x%x", rgb[0], rgb[1], rgb[2])
}
