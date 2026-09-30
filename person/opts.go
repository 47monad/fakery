// Package person generates fake person data: first/last names, prefixes,
// suffixes, and templated full names.
//
// It is the template migration for the three-tier API:
//
//	// Tier 2 — single-domain instance with its own seed.
//	p := person.New(person.WithLang("fa"), person.WithSeed(42))
//	p.Name(person.WithGender(person.Female))
//
//	// Tier 3 — zero setup package-level helpers.
//	person.FirstName()
//	person.LastName(person.WithLang("fa"))
//
// A seeded Faker is deterministic for single-goroutine use; concurrent
// use is safe but nondeterministic in call ordering.
package person

import (
	"github.com/47monad/fakery/fk"
)

// Gender selects the first-name pool a generator draws from.
//
// The zero value, GenderAny, is the default and picks randomly between
// the male and female pools instead of hardcoding one of them. Because it
// is a typed value, a misspelled gender cannot silently fall back to male
// at runtime; unknown values are normalized to GenderAny.
type Gender int

const (
	// GenderAny draws from both the male and female pools at random.
	GenderAny Gender = iota
	// Male draws only from the male first-name pool.
	Male
	// Female draws only from the female first-name pool.
	Female
)

// String returns the canonical lowercase name of the gender.
func (g Gender) String() string {
	switch g {
	case Male:
		return "male"
	case Female:
		return "female"
	default:
		return "any"
	}
}

// IsValid reports whether g is one of the defined Gender values.
func (g Gender) IsValid() bool {
	return g == GenderAny || g == Male || g == Female
}

// ParseGender maps a string to a Gender, tolerating case. Anything that is
// not "male" or "female" (including typos) becomes GenderAny — never a
// silent male default and never a panic.
func ParseGender(s string) Gender {
	switch s {
	case "male", "Male", "MALE", "m":
		return Male
	case "female", "Female", "FEMALE", "f":
		return Female
	default:
		return GenderAny
	}
}

// Options configures a Faker and is also accepted by every generator call.
//
// It embeds fk.Base so it carries the shared language + custom data
// override, which means the generic fk.WithLang[Options] helper works as
// well. Each option below is intentionally tiny.
type Options struct {
	fk.Base
	// Gender selects the first-name pool; the zero value GenderAny is
	// the default and never hardcodes male.
	Gender Gender
	// Seed is the deterministic RNG seed. HasSeed distinguishes an
	// explicit WithSeed(0) from the unseeded default.
	Seed    uint64
	HasSeed bool
}

// NameOpts, LastNameOpts, and FakerOpts are aliases of the unified Options
// type. A single options type is deliberate: it lets the same option value
// (notably person.WithLang, which cannot be generic over the target type)
// be passed to New and to every generator, while the aliases preserve the
// per-domain option names used elsewhere in the library.
type (
	// FakerOpts configures a Faker via New.
	FakerOpts = Options
	// NameOpts configures first-name and full-name generation.
	NameOpts = Options
	// LastNameOpts configures last-name generation.
	LastNameOpts = Options
)

// WithLang sets the language for any call. Invalid tags normalize to
// English (see fk.NormalizeLang) instead of panicking.
func WithLang(lang string) fk.Option[Options] {
	tag := fk.NormalizeLang(lang)
	return func(o *Options) {
		if o == nil {
			return
		}
		o.Lang = tag
	}
}

// WithSeed makes a Faker deterministic: the same seed always produces the
// same sequence for single-goroutine use. Without it, New seeds from
// crypto/rand.
func WithSeed(seed uint64) fk.Option[Options] {
	return func(o *Options) {
		if o == nil {
			return
		}
		o.Seed = seed
		o.HasSeed = true
	}
}

// WithGender selects the first-name pool for first-name and templated name
// generation. Values outside the defined set are normalized to GenderAny.
func WithGender(g Gender) fk.Option[Options] {
	if !g.IsValid() {
		g = GenderAny
	}
	return func(o *Options) {
		if o == nil {
			return
		}
		o.Gender = g
	}
}

// WithData overrides the embedded locale data. Accepted forms are Locale,
// *Locale, and *binder.Data[Locale]; anything else is ignored and the
// embedded English/localized data is used instead. It never panics.
func WithData(data any) fk.Option[Options] {
	return func(o *Options) {
		if o == nil {
			return
		}
		o.Data = data
	}
}

// WithPersonData is an alias for WithData, matching the name used in the
// package's issue sketch.
func WithPersonData(data any) fk.Option[Options] {
	return WithData(data)
}
