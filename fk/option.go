// Functional-options core for all faker domains.
//
// It replaces the hand-written Builder ritual in fk/fkopts with a small
// generic system:
//
//	type NameOpts struct { fk.Base; Gender Gender }
//	func WithGender(g Gender) fk.Option[NameOpts] {
//		return func(o *NameOpts) { o.Gender = g }
//	}
//	opts := fk.Build(WithGender(Male), fk.WithLang[NameOpts]("fa"))
//
// This file is additive: fk/fkopts keeps working until later issues
// migrate each domain.
package fk

import (
	"sync"

	"golang.org/x/text/language"
)

// Option is a functional option that mutates *T.
// A nil Option is valid and skipped by Build.
type Option[T any] func(*T)

// Base is embedded by every domain options struct to carry
// the shared language + custom data override.
type Base struct {
	Lang language.Tag
	Data any
}

// base promotes Base to the hasBase constraint so any struct
// embedding Base automatically satisfies hasBase.
func (b *Base) base() *Base { return b }

// hasBase restricts helpers to option structs that embed Base.
// It is satisfied by *T (not T) when T embeds Base by value, so
// helpers use runtime assertions instead of a [T hasBase] constraint
// to keep the call site fk.WithLang[NameOpts] ergonomic.
type hasBase interface {
	base() *Base
}

// defaulter is implemented by domain opts that provide defaults,
// e.g. func (o *NameOpts) Default() *NameOpts.
type defaulter[T any] interface {
	Default() *T
}

// plainDefaulter covers Default() without a return value.
type plainDefaulter interface {
	Default()
}

// supported Langs known to the library. This is a temporary fallback
// until internal/locales (#10) lands with per-domain matchers built
// from the embedded filesystem. en + fa covers all current datasets:
// every domain ships en.json, and lorem/person/phonenumber ship fa.json.
var supportedTags = []language.Tag{language.English, language.Persian}

var (
	matcherOnce sync.Once
	matcher     language.Matcher
)

func getMatcher() language.Matcher {
	matcherOnce.Do(func() {
		matcher = language.NewMatcher(supportedTags)
	})
	return matcher
}

// NormalizeLang parses s and matches it against supported locales.
// Unknown, malformed, or unmatched tags fall back to English.
// It never panics and never returns language.Und.
func NormalizeLang(s string) language.Tag {
	tag, err := language.Parse(s)
	if err != nil {
		return language.English
	}
	if tag == language.Und || tag.IsRoot() {
		return language.English
	}
	matched, _, conf := getMatcher().Match(tag)
	if conf == language.No {
		return language.English
	}
	// Strip extensions the matcher may add (e.g. "en-u-rg-uszzzz")
	// back to a clean tag when the base language matches.
	base, _ := matched.Base()
	if b, _ := tag.Base(); b.String() == base.String() {
		return tag
	}
	return matched
}

// WithLang returns an Option that sets the language for any opts
// struct embedding Base. Invalid tags normalize to English.
//
// NOTE: the issue sketch suggests [T hasBase], but Go method sets
// require base() on *Base, so only *T (not T) implements hasBase.
// Constraining T would reject value types like NameOpts at compile
// time. Using [T any] with a runtime assertion keeps the ergonomic
// call site fk.WithLang[NameOpts]("fa") working.
func WithLang[T any](lang string) Option[T] {
	tag := NormalizeLang(lang)
	return func(o *T) {
		if o == nil {
			return
		}
		if hb, ok := any(o).(hasBase); ok && hb != nil {
			if b := hb.base(); b != nil {
				b.Lang = tag
			}
		}
	}
}

// Build applies defaults (if T implements Default) then applies opts
// in order. Nil options are skipped. It never returns an error and
// never panics. Base.Lang is guaranteed to be a concrete tag
// (English if untouched).
func Build[T any](opts ...Option[T]) *T {
	o := new(T)
	applyDefaults(o)
	ensureLang(o)
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		opt(o)
	}
	ensureLang(o)
	return o
}

func applyDefaults[T any](o *T) {
	if d, ok := any(o).(defaulter[T]); ok && d != nil {
		d.Default()
		return
	}
	if d, ok := any(o).(plainDefaulter); ok && d != nil {
		d.Default()
	}
}

func ensureLang[T any](o *T) {
	hb, ok := any(o).(hasBase)
	if !ok || hb == nil {
		return
	}
	b := hb.base()
	if b == nil {
		return
	}
	if b.Lang == language.Und || b.Lang.IsRoot() {
		b.Lang = language.English
	}
}
