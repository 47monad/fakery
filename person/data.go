package person

import (
	"github.com/47monad/fakery/internal/binder"
	"github.com/47monad/fakery/internal/locales"
	"golang.org/x/text/language"
)

// domain is the locale directory name under resources/locales.
const domain = "person"

// Locale is the person dataset for a single language.
//
// Fields are optional per locale: a locale file may omit any key, in which
// case per-key fallback serves the English value (for example fa.json has
// no "suffix" or "nameWithMiddle").
type Locale struct {
	MaleFirstName   []string `json:"maleFirstName"`
	FemaleFirstName []string `json:"femaleFirstName"`
	LastName        []string `json:"lastName"`
	FullName        []string `json:"fullName"`
	NameWithMiddle  []string `json:"nameWithMiddle"`
	Prefix          []string `json:"prefix"`
	Suffix          []string `json:"suffix"`
}

// embeddedData returns the parse-once locale data for tag. It is cached by
// internal/locales and never returns nil: unknown locales degrade to
// English, and missing files degrade to empty (non-nil) data.
func embeddedData(tag language.Tag) *binder.Data[Locale] {
	return locales.Load[Locale](domain, tag)
}

// resolveData builds the data a Faker uses for a given language and custom
// data override.
//
// With no override it returns the embedded (fallback-aware) dataset. With a
// usable override it returns a dataset whose Default is the custom locale
// and whose Fallback is still the English one, so per-key fallback keeps
// working. Invalid or empty overrides fall back to the embedded dataset.
// It never panics.
func resolveData(tag language.Tag, raw any) *binder.Data[Locale] {
	embedded := embeddedData(tag)
	if raw == nil {
		return embedded
	}
	custom := customLocale(raw)
	if custom == nil || localeEmpty(custom) {
		return embedded
	}
	fallback := embedded.Fallback
	if fallback == nil {
		// Defensive: locales.Load guarantees a non-nil English fallback.
		fallback = embeddedData(language.English).Fallback
	}
	return &binder.Data[Locale]{Default: custom, Fallback: fallback}
}

// customLocale extracts a *Locale from the supported override forms.
// Unsupported values return nil.
func customLocale(raw any) *Locale {
	switch v := raw.(type) {
	case *Locale:
		return v
	case Locale:
		loc := v
		return &loc
	case *binder.Data[Locale]:
		if v != nil {
			return v.Default
		}
	}
	return nil
}

// localeEmpty reports whether every field of l is empty.
func localeEmpty(l *Locale) bool {
	if l == nil {
		return true
	}
	return len(l.MaleFirstName) == 0 &&
		len(l.FemaleFirstName) == 0 &&
		len(l.LastName) == 0 &&
		len(l.FullName) == 0 &&
		len(l.NameWithMiddle) == 0 &&
		len(l.Prefix) == 0 &&
		len(l.Suffix) == 0
}
