package fkcolor

import (
	"github.com/47monad/fakery/internal/fkloader"
	"github.com/47monad/fakery/internal/fktypes"
)

type ColorData struct {
	Name []string `json:"name"`
}

func NewColorData(opts *fktypes.FakeryConfig) *fktypes.ResultData[ColorData] {
	d, err := fkloader.FromJSON[ColorData](opts.FS, "color", opts.DefaultLang)
	if err != nil {
		panic(err)
	}
	return d
}

func keyName(d *ColorData) []string {
	return d.Name
}
