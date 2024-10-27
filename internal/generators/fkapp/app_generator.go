package fkapp

import (
	"fmt"

	"github.com/47monad/fakery/internal/fkgeneral"
	"github.com/47monad/fakery/internal/fklocaler"
	"github.com/47monad/fakery/internal/fktypes"
)

type AppGenerator struct {
	Name            func() string
	SemanticVersion func(...fktypes.SemanticVersionOption) string
}

func NewAppGenerator(opts *fktypes.FakeryConfig) AppGenerator {
	return AppGenerator{
		SemanticVersion: generateSemanticVersion(opts),
		Name:            generateAppName(opts),
	}
}

func generateSemanticVersion(opts *fktypes.FakeryConfig) func(...fktypes.SemanticVersionOption) string {
	return func(opts ...fktypes.SemanticVersionOption) string {
		_opts := fktypes.DefaultSemanticVersionOptions()

		for _, opt := range opts {
			opt(_opts)
		}

		return fmt.Sprintf("%d.%d.%d",
			fkgeneral.RandomInRange(_opts.Major.Min, _opts.Major.Max),
			fkgeneral.RandomInRange(_opts.Minor.Min, _opts.Minor.Max),
			fkgeneral.RandomInRange(_opts.Patch.Min, _opts.Patch.Max))
	}
}

func generateAppName(opts *fktypes.FakeryConfig) func() string {
	d := NewAppData(opts)
	return func() string {
		_, sampler := fklocaler.Localize(d, keyAppName)
		return sampler()
	}
}
