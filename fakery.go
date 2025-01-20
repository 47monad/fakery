package fakery

import (
	"github.com/47monad/fakery/fk/fklorem"
	"github.com/47monad/fakery/fk/fkopts"
	"golang.org/x/text/language"
)

var (
	fallbackLang = language.English
	defaultLang  = language.English
)

type Faker struct {
	lw *fkopts.LangWrapper
}

func New() *Faker {
	return &Faker{}
}

func (f *Faker) Word(opts ...fkopts.Builder[fkopts.WordOpts]) string {
	return fklorem.Word(opts...)
}

func (f *Faker) Words(opts ...fkopts.Builder[fkopts.WordsOpts]) []string {
	return fklorem.Words(opts...)
}

func (f *Faker) Sentence(opts ...fkopts.Builder[fkopts.SentenceOpts]) string {
	return fklorem.Sentence(opts...)
}
