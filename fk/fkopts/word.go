package fkopts

import (
	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/internal/binder"
)

type WordOpts struct {
	LW            *LangWrapper
	Data          *binder.Data[fkdata.Lorem]
	ExcludedWords []string
}

type WordBuilder struct {
	lw   *LangWrapper
	Opts []func(*WordOpts) error
}

func (b *WordBuilder) List() []func(*WordOpts) error {
	return append(b.Opts, func(wo *WordOpts) error {
		wo.LW = b.lw
		return nil
	})
}

func Word() *WordBuilder {
	return &WordBuilder{lw: &LangWrapper{}}
}

func (b *WordBuilder) SetLang(lang string) *WordBuilder {
	b.lw.UseLang(lang)
	return b
}

func (b *WordBuilder) SetData(data *binder.Data[fkdata.Lorem]) *WordBuilder {
	b.Opts = append(b.Opts, func(wo *WordOpts) error {
		wo.Data = data
		return nil
	})
	return b
}

func (b *WordBuilder) Exclude(word string, rest ...string) *WordBuilder {
	b.Opts = append(b.Opts, func(wo *WordOpts) error {
		wo.ExcludedWords = append([]string{word}, rest...)
		return nil
	})
	return b
}
