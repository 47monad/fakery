package fkdata

import (
	"github.com/47monad/fakery/internal/binder"
	"github.com/47monad/fakery/internal/locales"
	"golang.org/x/text/language"
)

type PhoneNumber struct {
	AreaCode     []string `json:"areaCode"`
	CountryCode  []string `json:"countryCode"`
	ExchangeCode []string `json:"exchangeCode"`
	Format       []string `json:"format"`
	MobileFormat []string `json:"mobileFormat"`
}

func NewPhoneNumber(lang language.Tag) *binder.Data[PhoneNumber] {
	return locales.Load[PhoneNumber]("phonenumber", lang)
}
