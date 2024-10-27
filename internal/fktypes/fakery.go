package fktypes

import (
	"golang.org/x/text/language"
	"io/fs"
)

type FakeryConfig struct {
	FS           fs.FS
	DefaultLang  language.Tag
	FallbackLang language.Tag
}

type FakeryOption func(*FakeryConfig)

func WithDefaultLang(fs fs.FS) FakeryOption {
	return func(o *FakeryConfig) {
		o.FS = fs
	}
}

type ResultData[K any] struct {
	Default  *K
	Fallback *K
}
