package fkinternet

import (
	"fmt"

	"github.com/47monad/fakery/internal/fkgeneral"
	"github.com/47monad/fakery/internal/fklocaler"
	"github.com/47monad/fakery/internal/fktypes"
)

type InternetGenerator struct {
	IPv4Address   func() string
	IPv6Address   func() string
	MacAddress    func() string
	Browser       func() string
	BrowserEngine func() string
}

func NewInternetGenerator(opts *fktypes.FakeryConfig) InternetGenerator {
	return InternetGenerator{
		IPv4Address:   generateIPv4Address(opts),
		IPv6Address:   generateIPv6Address(opts),
		MacAddress:    generateMacAddress(opts),
		Browser:       generateBrowser(opts),
		BrowserEngine: generateBrowserEngine(opts),
	}
}

func generateIPv4Address(opts *fktypes.FakeryConfig) func() string {
	return func() string {
		num := func() int { return fkgeneral.RandomInt(256) }

		return fmt.Sprintf("%d.%d.%d.%d", num(), num(), num(), num())
	}
}

func generateIPv6Address(opts *fktypes.FakeryConfig) func() string {
	return func() string {
		num := func() int { return fkgeneral.RandomInt(65536) }

		return fmt.Sprintf("%x:%x:%x:%x:%x:%x:%x:%x", num(), num(), num(), num(), num(), num(), num(), num())
	}
}

func generateMacAddress(opts *fktypes.FakeryConfig) func() string {
	return func() string {
		num := func() int { return fkgeneral.RandomInt(255) }

		return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", num(), num(), num(), num(), num(), num())
	}
}

func generateBrowser(opts *fktypes.FakeryConfig) func() string {
	d := NewInternetData(opts)
	return func() string {
		_, sampler := fklocaler.Localize(d, keyBrowser)
		return sampler()
	}
}

func generateBrowserEngine(opts *fktypes.FakeryConfig) func() string {
	d := NewInternetData(opts)
	return func() string {
		_, sampler := fklocaler.Localize(d, keyBrowserEngine)
		return sampler()
	}
}
