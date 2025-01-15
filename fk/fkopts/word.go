package fkopts

import (
	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/internal/binder"
	"golang.org/x/text/language"
)

type Lang struct {
	Tag language.Tag
}

func NewLang(lang string) (Lang, error) {
	tag, err := language.Parse(lang)
	if err != nil {
		return Lang{}, err
	}
	return Lang{Tag: tag}, nil
}

type WordOpts struct {
	Lang
	Data          *binder.Data[fkdata.Lorem]
	ExcludedWords []string
}

type WordBuilder struct {
	Opts []func(*WordOpts) error
}

func (b *WordBuilder) List() []func(*WordOpts) error {
	return b.Opts
}

func Word() *WordBuilder {
	return &WordBuilder{}
}

func (b *WordBuilder) SetLang(lang string) *WordBuilder {
	b.Opts = append(b.Opts, func(wo *WordOpts) error {
		l, err := NewLang(lang)
		if err != nil {
			return err
		}
		wo.Lang = l

		return nil
	})

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
