package fkopts

import (
	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/internal/binder"
)

type FakerOpts struct {
	Lang
	Data          *binder.Data[fkdata.Lorem]
	ExcludedWords []string
}

type FakerBuilder struct {
	Opts []func(*FakerOpts) error
}

func (b *FakerBuilder) List() []func(*FakerOpts) error {
	return b.Opts
}

func Faker() *FakerBuilder {
	return &FakerBuilder{}
}

func (b *FakerBuilder) SetLang(lang string) *FakerBuilder {
	b.Opts = append(b.Opts, func(wo *FakerOpts) error {
		l, err := NewLang(lang)
		if err != nil {
			return err
		}
		wo.Lang = l

		return nil
	})

	return b
}

func (b *FakerBuilder) SetData(data *binder.Data[fkdata.Lorem]) *FakerBuilder {
	b.Opts = append(b.Opts, func(wo *FakerOpts) error {
		wo.Data = data

		return nil
	})

	return b
}
