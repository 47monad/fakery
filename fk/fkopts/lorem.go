package fkopts

import (
	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/internal/binder"
)

// Word Options
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

// Words Options
type WordsOpts struct {
	LW    *LangWrapper
	Data  *binder.Data[fkdata.Lorem]
	Count int
}

type WordsBuilder struct {
	lw   *LangWrapper
	Opts []func(*WordsOpts) error
}

func (b *WordsBuilder) List() []func(*WordsOpts) error {
	return append(b.Opts, func(wo *WordsOpts) error {
		wo.LW = b.lw
		return nil
	})
}

func Words() *WordsBuilder {
	return &WordsBuilder{lw: &LangWrapper{}}
}

func (b *WordsBuilder) SetLang(lang string) *WordsBuilder {
	b.lw.UseLang(lang)
	return b
}

func (b *WordsBuilder) SetData(data *binder.Data[fkdata.Lorem]) *WordsBuilder {
	b.Opts = append(b.Opts, func(op *WordsOpts) error {
		op.Data = data
		return nil
	})
	return b
}

func (b *WordsBuilder) SetCount(count int) *WordsBuilder {
	b.Opts = append(b.Opts, func(op *WordsOpts) error {
		op.Count = count
		return nil
	})
	return b
}

// Sentence Options
type SentenceOpts struct {
	LW        *LangWrapper
	Data      *binder.Data[fkdata.Lorem]
	WordCount int
}

type SentenceBuilder struct {
	lw   *LangWrapper
	Opts []func(*SentenceOpts) error
}

func (b *SentenceBuilder) List() []func(*SentenceOpts) error {
	return append(b.Opts, func(wo *SentenceOpts) error {
		wo.LW = b.lw
		return nil
	})
}

func Sentence() *SentenceBuilder {
	return &SentenceBuilder{lw: &LangWrapper{}}
}

func (b *SentenceBuilder) SetLang(lang string) *SentenceBuilder {
	b.lw.UseLang(lang)
	return b
}

func (b *SentenceBuilder) SetData(data *binder.Data[fkdata.Lorem]) *SentenceBuilder {
	b.Opts = append(b.Opts, func(op *SentenceOpts) error {
		op.Data = data
		return nil
	})
	return b
}

func (b *SentenceBuilder) SetWordCount(count int) *SentenceBuilder {
	b.Opts = append(b.Opts, func(op *SentenceOpts) error {
		op.WordCount = count
		return nil
	})
	return b
}
