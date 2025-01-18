package fkopts

import (
	"github.com/47monad/fakery/internal/binder"
)

type FakerOpts[K any] struct {
	LW   *LangWrapper
	Data *binder.Data[K]
}

type FakerBuilder[K any] struct {
	lw   *LangWrapper
	Opts []func(*FakerOpts[K]) error
}

func (b *FakerBuilder[K]) List() []func(*FakerOpts[K]) error {
	return append(b.Opts, func(f *FakerOpts[K]) error {
		f.LW = b.lw
		return nil
	})
}

func Faker[K any]() *FakerBuilder[K] {
	return &FakerBuilder[K]{lw: &LangWrapper{}}
}

func (b *FakerBuilder[K]) SetLang(lang string) *FakerBuilder[K] {
	b.lw.UseLang(lang)
	return b
}

func (b *FakerBuilder[K]) SetData(data *binder.Data[K]) *FakerBuilder[K] {
	b.Opts = append(b.Opts, func(wo *FakerOpts[K]) error {
		wo.Data = data

		return nil
	})

	return b
}
