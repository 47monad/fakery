package fkinternet

import (
	"github.com/47monad/fakery/internal/fkloader"
	"github.com/47monad/fakery/internal/fktypes"
)

type InternetData struct {
	Browser          []string `json:"browser"`
	BrowserEngine    []string `json:"browserEngine"`
	DomainSuffix     []string `json:"domainSuffix"`
	SafeDomainSuffix []string `json:"safeDomainSuffix"`
	Slug             []string `json:"slug"`
}

func NewInternetData(opts *fktypes.FakeryConfig) *fktypes.ResultData[InternetData] {
	d, err := fkloader.FromJSON[InternetData](opts.FS, "internet", opts.DefaultLang)
	if err != nil {
		panic(err)
	}
	return d
}

func keyBrowser(d *InternetData) []string {
	return d.Browser
}

func keyBrowserEngine(d *InternetData) []string {
	return d.BrowserEngine
}

func keyDomainSuffix(d *InternetData) []string {
	return d.DomainSuffix
}

func keySafeDomainSuffix(d *InternetData) []string {
	return d.SafeDomainSuffix
}

func keySlug(d *InternetData) []string {
	return d.Slug
}
