package fkdata

import (
	"github.com/47monad/fakery/internal/binder"
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
	d, err := binder.JSON[Device]("device", lang)
	if err != nil {
		panic(err)
	}
	return d
}
