package fkdevice

import (
	"github.com/47monad/fakery/internal/fkloader"
	"github.com/47monad/fakery/internal/fktypes"
)

type DeviceData struct {
	Manufacturer []string `json:"manufacturer"`
	ModelName    []string `json:"modelName"`
	SerialNumber []string `json:"serial"`
	Platform     []string `json:"platform"`
	Type         []string `json:"type"`
}

func NewDeviceData(opts *fktypes.FakeryConfig) *fktypes.ResultData[DeviceData] {
	d, err := fkloader.FromJSON[DeviceData](opts.FS, "device", opts.DefaultLang)
	if err != nil {
		panic(err)
	}
	return d
}

func keyDeviceModel(data *DeviceData) []string {
	return data.ModelName
}

func keyDeviceManufacturer(d *DeviceData) []string {
	return d.Manufacturer
}

func keyDevicePlatform(d *DeviceData) []string {
	return d.Platform
}

func keyDeviceType(d *DeviceData) []string {
	return d.Type
}

func keyDeviceSerialNumber(d *DeviceData) []string {
	return d.SerialNumber
}
