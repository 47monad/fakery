package fakery

import (
	"embed"

	"github.com/47monad/fakery/internal/fkgeneral"
	"github.com/47monad/fakery/internal/fktypes"
	"github.com/47monad/fakery/internal/generators/fkapp"
	"github.com/47monad/fakery/internal/generators/fkcolor"
	"github.com/47monad/fakery/internal/generators/fkcompany"
	"github.com/47monad/fakery/internal/generators/fkdevice"
	"github.com/47monad/fakery/internal/generators/fkinternet"
	"github.com/47monad/fakery/internal/generators/fknumber"
	"github.com/47monad/fakery/internal/generators/fkperson"
	"github.com/47monad/fakery/internal/generators/fkphonenumber"
	"golang.org/x/text/language"
)

//go:embed resources/locales/**/*.json
var resourcesFS embed.FS

var (
	fallbackLang = language.English
	defaultLang  = language.English
)

func SetDefaultLang(lang language.Tag) {
	defaultLang = lang
}

func prepareConfig() *fktypes.FakeryConfig {
	return &fktypes.FakeryConfig{
		FS:           resourcesFS,
		DefaultLang:  defaultLang,
		FallbackLang: fallbackLang,
	}
}

func Person() fkperson.PersonGenerator {
	return fkperson.NewPersonGenerator(prepareConfig())
}

func Device() fkdevice.DeviceGenerator {
	return fkdevice.NewDeviceGenerator(prepareConfig())
}

func App() fkapp.AppGenerator {
	return fkapp.NewAppGenerator(prepareConfig())
}

func Internet() fkinternet.InternetGenerator {
	return fkinternet.NewInternetGenerator(prepareConfig())
}

func Phone() fkphonenumber.PhoneNumberGenerator {
	return fkphonenumber.NewPhoneNumberGenerator(prepareConfig())
}

func Company() fkcompany.CompanyGenerator {
	return fkcompany.NewCompanyGenerator(prepareConfig())
}

func Number() fknumber.NumberGenerator {
	return fknumber.NewNumberGenerator(prepareConfig())
}

func Color() fkcolor.ColorGenerator {
	return fkcolor.NewColorGenerator(prepareConfig())
}

func RandomElement[K any](slice []K) K {
	return fkgeneral.RandomElement[K](slice)
}
