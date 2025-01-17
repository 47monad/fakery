package fkdata

import (
	"github.com/47monad/fakery/internal/binder"
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
	d, err := binder.JSON[PhoneNumber]("phonenumber", lang)
	if err != nil {
		panic(err)
	}
	return d
}
