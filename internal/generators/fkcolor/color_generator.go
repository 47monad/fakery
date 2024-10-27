package fkcolor

import (
	"fmt"

	"github.com/47monad/fakery/internal/fkgeneral"
	"github.com/47monad/fakery/internal/fklocaler"
	"github.com/47monad/fakery/internal/fktypes"
)

type ColorGenerator struct {
	Name     func() string
	RGBColor func() [3]int
	HexColor func() string
}

func NewColorGenerator(opts *fktypes.FakeryConfig) ColorGenerator {
	return ColorGenerator{
		Name:     generateName(opts),
		RGBColor: generateRGBColor,
		HexColor: generateHexColor,
	}
}

func generateName(opts *fktypes.FakeryConfig) func() string {
	return func() string {
		d := NewColorData(opts)
		_, sampler := fklocaler.Localize(d, keyName)
		return sampler()
	}
}

func generateRGBColor() [3]int {
	num := func() int { return fkgeneral.RandomInRange(0, 255) }

	return [3]int{num(), num(), num()}
}

func generateHexColor() string {
	rgb := generateRGBColor()

	return fmt.Sprintf("#%x%x%x", rgb[0], rgb[1], rgb[2])
}
