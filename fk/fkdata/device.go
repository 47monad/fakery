package fkdata

import (
	"github.com/47monad/fakery/internal/binder"
	"github.com/47monad/fakery/internal/locales"
	"golang.org/x/text/language"
)

type Device struct {
	Manufacturer []string `json:"manufacturer"`
	ModelName    []string `json:"modelName"`
	SerialNumber []string `json:"serial"`
	Platform     []string `json:"platform"`
	Type         []string `json:"type"`
}

func NewDevice(lang language.Tag) *binder.Data[Device] {
	return locales.Load[Device]("device", lang)
}
