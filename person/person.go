package person

import (
	"text/template"

	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/internal/binder"
	"github.com/47monad/fakery/internal/rng"
	"github.com/47monad/fakery/internal/templater"
	"golang.org/x/text/language"
)

// Faker is a single-domain person faker.
//
// Construct one with New when you want an independent, optionally seeded
// stream:
//
//	p := person.New(person.WithLang("fa"), person.WithSeed(42))
//	p.Name(person.WithGender(person.Female))
//
// A seeded Faker is deterministic for single-goroutine use. Concurrent use
// is safe (the underlying RNG holds a mutex) but nondeterministic in call
// ordering. The zero value is usable and behaves like the package-level
// default (English, unseeded); a nil *Faker is tolerated too.
type Faker struct {
	opts Options
	data *binder.Data[Locale]
	r    *rng.RNG
}

// New returns a Faker. With no options it uses English and an unseeded RNG;
// with WithSeed it is deterministic; with WithLang and WithData it selects
// locale data. It never panics, even on invalid language or custom data.
func New(opts ...fk.Option[Options]) *Faker {
	o := fk.Build(opts...)
	f := &Faker{opts: *o}
	if o.HasSeed {
		f.r = rng.New(o.Seed)
	} else {
		f.r = rng.Default()
	}
	f.data = resolveData(o.Lang, o.Data)
	return f
}

// defaultFaker backs every tier-3 package-level function. It uses an
// unseeded RNG, so tier-3 output is non-deterministic by design.
var defaultFaker = New()

// resolve applies opts on top of the Faker's own options, normalizes the
// language, and normalizes an out-of-range Gender. It never panics.
func (f *Faker) resolve(opts []fk.Option[Options]) Options {
	var o Options
	if f != nil {
		o = f.opts
	}
	for _, opt := range opts {
		if opt != nil {
			opt(&o)
		}
	}
	if o.Lang == language.Und || o.Lang.IsRoot() {
		o.Lang = language.English
	}
	if !o.Gender.IsValid() {
		o.Gender = GenderAny
	}
	return o
}

// dataFor returns the data matching resolved options, reusing the Faker's
// cached dataset when nothing locale-related changed.
func (f *Faker) dataFor(o Options) *binder.Data[Locale] {
	if f != nil && f.data != nil && o.Lang == f.opts.Lang && o.Data == nil {
		return f.data
	}
	return resolveData(o.Lang, o.Data)
}

// rngOrDefault guards against a nil Faker or nil RNG so methods never panic
// on a nil receiver; callers get the default (unseeded) stream.
func (f *Faker) rngOrDefault() *rng.RNG {
	if f != nil && f.r != nil {
		return f.r
	}
	return defaultFaker.r
}

// FirstName returns a first name for the configured language and gender.
// With GenderAny it draws from both the male and female pools.
func (f *Faker) FirstName(opts ...fk.Option[Options]) string {
	o := f.resolve(opts)
	return firstNameFrom(f.dataFor(o), f.rngOrDefault(), o.Gender)
}

// LastName returns a last name for the configured language.
func (f *Faker) LastName(opts ...fk.Option[Options]) string {
	o := f.resolve(opts)
	return lastNameFrom(f.dataFor(o), f.rngOrDefault())
}

// Name returns a full name rendered from a locale full-name template.
func (f *Faker) Name(opts ...fk.Option[Options]) string {
	o := f.resolve(opts)
	d := f.dataFor(o)
	tmpl := pick(d, f.rngOrDefault(), func(l *Locale) []string { return l.FullName })
	return f.exec(d, o.Gender, tmpl)
}

// NameWithMiddle returns a name rendered from a locale name-with-middle
// template. Locales without one (e.g. fa) fall back per key to English.
func (f *Faker) NameWithMiddle(opts ...fk.Option[Options]) string {
	o := f.resolve(opts)
	d := f.dataFor(o)
	tmpl := pick(d, f.rngOrDefault(), func(l *Locale) []string { return l.NameWithMiddle })
	return f.exec(d, o.Gender, tmpl)
}

// Prefix returns an honorific such as "Dr." or "Prof.".
func (f *Faker) Prefix(opts ...fk.Option[Options]) string {
	o := f.resolve(opts)
	return prefixFrom(f.dataFor(o), f.rngOrDefault())
}

// Suffix returns a name suffix such as "Jr." or "PhD".
func (f *Faker) Suffix(opts ...fk.Option[Options]) string {
	o := f.resolve(opts)
	return suffixFrom(f.dataFor(o), f.rngOrDefault())
}

// exec renders tmpl with the faker's generators bound to d. Parse and
// execute failures degrade to the raw template string instead of panicking.
func (f *Faker) exec(d *binder.Data[Locale], g Gender, tmpl string) string {
	if tmpl == "" {
		return ""
	}
	r := f.rngOrDefault()
	funcs := template.FuncMap{
		"prefix":    func() string { return prefixFrom(d, r) },
		"suffix":    func() string { return suffixFrom(d, r) },
		"firstName": func() string { return firstNameFrom(d, r, g) },
		"lastName":  func() string { return lastNameFrom(d, r) },
	}
	out, err := templater.Exec("person", tmpl, funcs)
	if err != nil {
		return tmpl
	}
	return out
}

// pick chooses one value from the first non-empty pool among the locale's
// Default then Fallback, honoring per-key fallback. Empty input yields "".
func pick(d *binder.Data[Locale], r *rng.RNG, field func(*Locale) []string) string {
	if d == nil {
		return ""
	}
	if d.Default != nil {
		if pool := field(d.Default); len(pool) > 0 {
			return pool[r.IntN(len(pool))]
		}
	}
	if d.Fallback != nil {
		if pool := field(d.Fallback); len(pool) > 0 {
			return pool[r.IntN(len(pool))]
		}
	}
	return ""
}

func firstNameFrom(d *binder.Data[Locale], r *rng.RNG, g Gender) string {
	switch g {
	case Male:
		return pick(d, r, func(l *Locale) []string { return l.MaleFirstName })
	case Female:
		return pick(d, r, func(l *Locale) []string { return l.FemaleFirstName })
	default:
		// GenderAny: pick a pool at random, then a name from it.
		if r.Bool() {
			return pick(d, r, func(l *Locale) []string { return l.MaleFirstName })
		}
		return pick(d, r, func(l *Locale) []string { return l.FemaleFirstName })
	}
}

func lastNameFrom(d *binder.Data[Locale], r *rng.RNG) string {
	return pick(d, r, func(l *Locale) []string { return l.LastName })
}

func prefixFrom(d *binder.Data[Locale], r *rng.RNG) string {
	return pick(d, r, func(l *Locale) []string { return l.Prefix })
}

func suffixFrom(d *binder.Data[Locale], r *rng.RNG) string {
	return pick(d, r, func(l *Locale) []string { return l.Suffix })
}
