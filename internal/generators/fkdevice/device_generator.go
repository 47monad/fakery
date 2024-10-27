package fkdevice

import (
	"github.com/47monad/fakery/internal/fklocaler"
	"github.com/47monad/fakery/internal/fktypes"
)

type DeviceGenerator struct {
	Manufacturer func() string
	Type         func() string
	Model        func() string
	SerialNumber func() string
	Platform     func() string
}

func NewDeviceGenerator(opts *fktypes.FakeryConfig) DeviceGenerator {
	return DeviceGenerator{
		Manufacturer: generateDeviceManufacturer(opts),
		Type:         generateDeviceType(opts),
		Model:        generateDeviceModel(opts),
		SerialNumber: generateDeviceSerialNumber(opts),
		Platform:     generateDevicePlatform(opts),
	}
}
func generateDeviceModel(opts *fktypes.FakeryConfig) func() string {
	data := NewDeviceData(opts)
	return func() string {
		_, sampler := fklocaler.Localize(data, keyDeviceModel)
		return sampler()
	}
}

func generateDeviceManufacturer(opts *fktypes.FakeryConfig) func() string {
	data := NewDeviceData(opts)
	return func() string {
		_, sampler := fklocaler.Localize(data, keyDeviceManufacturer)
		return sampler()
	}
}

func generateDevicePlatform(opts *fktypes.FakeryConfig) func() string {
	data := NewDeviceData(opts)
	return func() string {
		_, sampler := fklocaler.Localize(data, keyDevicePlatform)
		return sampler()
	}
}

func generateDeviceType(opts *fktypes.FakeryConfig) func() string {
	data := NewDeviceData(opts)
	return func() string {
		_, sampler := fklocaler.Localize(data, keyDeviceType)
		return sampler()
	}
}

func generateDeviceSerialNumber(opts *fktypes.FakeryConfig) func() string {
	data := NewDeviceData(opts)
	return func() string {
		_, sampler := fklocaler.Localize(data, keyDeviceSerialNumber)
		return sampler()
	}
}
