package fkinternet

import (
	"fmt"
	"strings"

	"github.com/47monad/fakery/internal/fkgeneral"
	"github.com/47monad/fakery/internal/fklocaler"
	"github.com/47monad/fakery/internal/fktypes"
	"github.com/47monad/fakery/internal/generators/fkcompany"
	"github.com/47monad/fakery/internal/generators/fkperson"
	"golang.org/x/text/language"
)

type InternetGenerator struct {
	IPv4Address   func() string
	IPv6Address   func() string
	MacAddress    func() string
	Browser       func() string
	BrowserEngine func() string
	DomainSuffix  func(bool) string
	Slug          func() string
	// TODO. add arguments
	DomainName func() string
	Email      func() string
	// TODO. add arguments to modify the generation behaviour
	Username func() string
	// TODO.
	Password func() string
}

func NewInternetGenerator(opts *fktypes.FakeryConfig) InternetGenerator {
	return InternetGenerator{
		IPv4Address:   generateIPv4Address(opts),
		IPv6Address:   generateIPv6Address(opts),
		MacAddress:    generateMacAddress(opts),
		Browser:       generateBrowser(opts),
		BrowserEngine: generateBrowserEngine(opts),
		DomainSuffix:  generateDomainSuffix(opts),
		DomainName:    generateDomainName(opts),
		Email:         generateEmail(opts),
		Username:      generateUsername(opts),
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

func generateDomainSuffix(opts *fktypes.FakeryConfig) func(bool) string {
	d := NewInternetData(opts)
	return func(isSafe bool) string {
		keyer := keyDomainSuffix
		if isSafe {
			keyer = keySafeDomainSuffix
		}

		_, sampler := fklocaler.Localize(d, keyer)
		return sampler()
	}
}

func generateDomainName(opts *fktypes.FakeryConfig) func() string {
	return func() string {
		companyName := strings.Replace(fkcompany.NewCompanyGenerator(opts).Buzzword(), " ", "", -1)
		suffix := generateDomainSuffix(opts)(true)
		return fmt.Sprintf("%s.%s", strings.ToLower(companyName), suffix)
	}
}

func generateEmail(opts *fktypes.FakeryConfig) func() string {
	return func() string {
		opts := &fktypes.FakeryConfig{
			FS:           opts.FS,
			FallbackLang: language.English,
			DefaultLang:  language.English,
		}
		domain := generateDomainName(opts)()
		firstName := fkperson.NewPersonGenerator(opts).FirstName("male")
		lastName := fkperson.NewPersonGenerator(opts).LastName()

		return fmt.Sprintf("%s.%s@%s", strings.ToLower(firstName), strings.ToLower(lastName), domain)
	}
}

func generateUsername(opts *fktypes.FakeryConfig) func() string {
	return func() string {
		opts := &fktypes.FakeryConfig{
			FS:           opts.FS,
			FallbackLang: language.English,
			DefaultLang:  language.English,
		}
		lastName := fkperson.NewPersonGenerator(opts).LastName()

		return fmt.Sprintf("%s%d", strings.ToLower(lastName), fkgeneral.RandomInt(9999))
	}
}
