package fkphonenumber

import (
	"fmt"

	"github.com/47monad/fakery/internal/fklocaler"
	"github.com/47monad/fakery/internal/fktemplater"
	"github.com/47monad/fakery/internal/fktypes"
	"github.com/47monad/fakery/internal/generators/fknumber"
)

type PhoneNumberGenerator struct {
	AreaCode     func() string
	CountryCode  func() string
	ExchangeCode func() string
	Extension    func(uint32) string
	PhoneNumber  func() string
	MobileNumber func() string
}

func NewPhoneNumberGenerator(opts *fktypes.FakeryConfig) PhoneNumberGenerator {
	return PhoneNumberGenerator{
		AreaCode:     generatePhoneNumberAreaCode(opts),
		CountryCode:  generatePhoneNumberCountryCode(opts),
		ExchangeCode: generatePhoneNumberExchangeCode(opts),
		Extension:    generatePhoneNumberExtension(opts),
		PhoneNumber:  generatePhoneNumber(opts),
		MobileNumber: generateMobileNumber(opts),
	}
}

func generatePhoneNumberAreaCode(opts *fktypes.FakeryConfig) func() string {
	d := newPhoneNumberData(opts)
	return func() string {
		_, sampler := fklocaler.Localize(d, keyPhoneNumberAreaCode)
		return sampler()
	}
}

func generatePhoneNumberCountryCode(opts *fktypes.FakeryConfig) func() string {
	d := newPhoneNumberData(opts)
	return func() string {
		_, sampler := fklocaler.Localize(d, keyPhoneNumberCountryCode)
		return "+" + sampler()
	}
}

func generatePhoneNumberExchangeCode(opts *fktypes.FakeryConfig) func() string {
	d := newPhoneNumberData(opts)
	return func() string {
		_, sampler := fklocaler.Localize(d, keyPhoneNumberExchangeCode)
		return sampler()
	}
}

func generatePhoneNumberExtension(opts *fktypes.FakeryConfig) func(uint32) string {
	g := fknumber.NewNumberGenerator(opts)
	return func(length uint32) string {
		return fmt.Sprintf("%d", g.Number(length))
	}
}

func generatePhoneNumber(opts *fktypes.FakeryConfig) func() string {
	d := newPhoneNumberData(opts)
	return func() string {
		data, sampler := fklocaler.Localize(d, keyPhoneNumberFormat)
		return fktemplater.MustExecTemplate("phoneNumber", sampler(), PhoneNumberGFunc(data))
	}
}

func generateMobileNumber(opts *fktypes.FakeryConfig) func() string {
	d := newPhoneNumberData(opts)
	return func() string {
		data, sampler := fklocaler.Localize(d, keyMobileNumberFormat)
		return fktemplater.MustExecTemplate("mobileNumber", sampler(), PhoneNumberGFunc(data))
	}
}
