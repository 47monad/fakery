package fkphonenumber

import (
	"github.com/47monad/fakery/internal/fkloader"
	"github.com/47monad/fakery/internal/fktypes"
)

type PhoneNumberData struct {
	AreaCode     []string `json:"areaCode"`
	CountryCode  []string `json:"countryCode"`
	ExchangeCode []string `json:"exchangeCode"`
	Format       []string `json:"format"`
	MobileFormat []string `json:"mobileFormat"`
}

func newPhoneNumberData(opts *fktypes.FakeryConfig) *fktypes.ResultData[PhoneNumberData] {
	d, err := fkloader.FromJSON[PhoneNumberData](opts.FS, "phonenumber", opts.DefaultLang)
	if err != nil {
		panic(err)
	}
	return d
}

func keyPhoneNumberAreaCode(d *PhoneNumberData) []string {
	return d.AreaCode
}

func keyPhoneNumberCountryCode(d *PhoneNumberData) []string {
	return d.CountryCode
}

func keyPhoneNumberExchangeCode(d *PhoneNumberData) []string {
	return d.ExchangeCode
}

func keyPhoneNumberFormat(d *PhoneNumberData) []string {
	return d.Format
}

func keyMobileNumberFormat(d *PhoneNumberData) []string {
	return d.MobileFormat
}
