package fkopts

import (
	"github.com/47monad/fakery/internal/binder"
)

type FakerOpts[K any] struct {
	Lang
	Data          *binder.Data[K]
	ExcludedWords []string
}

type FakerBuilder[K any] struct {
	Opts []func(*FakerOpts[K]) error
}

func (b *FakerBuilder[K]) List() []func(*FakerOpts[K]) error {
	return b.Opts
}

func Faker[K any]() *FakerBuilder[K] {
	return &FakerBuilder[K]{}
}

func (b *FakerBuilder[K]) SetLang(lang string) *FakerBuilder[K] {
	b.Opts = append(b.Opts, func(wo *FakerOpts[K]) error {
		l, err := NewLang(lang)
		if err != nil {
			return err
		}
		wo.Lang = l

		return nil
	})

	return b
}

func (b *FakerBuilder[K]) SetData(data *binder.Data[K]) *FakerBuilder[K] {
	b.Opts = append(b.Opts, func(wo *FakerOpts[K]) error {
		wo.Data = data

		return nil
	})

	return b
}
