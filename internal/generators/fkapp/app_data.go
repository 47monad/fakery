package fkapp

import (
	"github.com/47monad/fakery/internal/fkloader"
	"github.com/47monad/fakery/internal/fktypes"
)

type AppData struct {
	Name []string `json:"name"`
}

func NewAppData(opts *fktypes.FakeryConfig) *fktypes.ResultData[AppData] {
	d, err := fkloader.FromJSON[AppData](opts.FS, "app", opts.DefaultLang)
	if err != nil {
		panic(err)
	}
	return d
}

func keyAppName(d *AppData) []string {
	return d.Name
}
