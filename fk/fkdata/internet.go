package fkdata

import (
	"github.com/47monad/fakery/internal/binder"
	"golang.org/x/text/language"
)

type Internet struct {
	Browser          []string `json:"browser"`
	BrowserEngine    []string `json:"browserEngine"`
	DomainSuffix     []string `json:"domainSuffix"`
	SafeDomainSuffix []string `json:"safeDomainSuffix"`
	Slug             []string `json:"slug"`
}

func NewInternet(lang language.Tag) *binder.Data[Internet] {
	d, err := binder.JSON[Internet]("internet", lang)
	if err != nil {
		panic(err)
	}
	return d
}
