package fklorem

import (
	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/fk/fkopts"
)

func Word(opts ...fkopts.Builder[fkopts.WordOpts]) string {
	_opts, _ := fkopts.Build(opts...)
	ctx := fk.NewContext(_opts.LW.GetFirstLang(), _opts.Data).SetBinder(fkdata.NewLorem)
	return word(ctx, _opts.ExcludedWords)
}

func Words(opts ...fkopts.Builder[fkopts.WordsOpts]) []string {
	_opts, _ := fkopts.Build(opts...)
	ctx := fk.NewContext(_opts.LW.GetFirstLang(), _opts.Data).SetBinder(fkdata.NewLorem)
	return words(ctx, _opts.Count)
}

func Sentence(opts ...fkopts.Builder[fkopts.SentenceOpts]) string {
	_opts, _ := fkopts.Build(opts...)
	ctx := fk.NewContext(_opts.LW.GetFirstLang(), _opts.Data).SetBinder(fkdata.NewLorem)
	return sentence(ctx, _opts.WordCount)
}
