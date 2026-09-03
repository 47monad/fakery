package fk_test

import (
	"testing"

	"github.com/47monad/fakery/fk"
	"golang.org/x/text/language"
)

// exampleDomainOpts is the compile-time check from the issue:
// an example domain opts struct + 2 options in < 20 lines.
type Gender string

const ( //nolint:revive // example domain enum
	Male   Gender = "male"
	Female Gender = "female"
)

type NameOpts struct {
	fk.Base
	Gender Gender
}

func (o *NameOpts) Default() *NameOpts { o.Gender = Male; return o }
func WithGender(g Gender) fk.Option[NameOpts] {
	return func(o *NameOpts) { o.Gender = g }
}
func WithNickname(n string) fk.Option[NameOpts] {
	return func(o *NameOpts) { o.Data = n }
}

func TestBuildZeroOptionsReturnsDefaults(t *testing.T) {
	got := fk.Build[NameOpts]()
	if got.Gender != Male {
		t.Fatalf("expected default gender %q, got %q", Male, got.Gender)
	}
	if got.Lang != language.English {
		t.Fatalf("expected default lang English, got %q", got.Lang.String())
	}
}

func TestLaterOptionsOverrideEarlier(t *testing.T) {
	got := fk.Build(WithGender(Male), WithGender(Female))
	if got.Gender != Female {
		t.Fatalf("expected later option to win, got %q", got.Gender)
	}
}

func TestWithLangInvalidNormalizesToEnglish(t *testing.T) {
	for _, s := range []string{"not-a-language", "bogus", "", "xx", "fr"} {
		got := fk.Build(fk.WithLang[NameOpts](s))
		if got.Lang != language.English {
			t.Fatalf("WithLang(%q) = %q, want English", s, got.Lang.String())
		}
	}
}

func TestWithLangValid(t *testing.T) {
	got := fk.Build(fk.WithLang[NameOpts]("fa"))
	if got.Lang.String() != "fa" && got.Base.Lang != language.Persian {
		t.Fatalf("WithLang(fa) = %q, want fa", got.Lang.String())
	}
	// Region variant resolves sanely without panic.
	gotUS := fk.Build(fk.WithLang[NameOpts]("en-US"))
	if gotUS.Lang.String() != "en-US" && gotUS.Lang != language.English {
		t.Fatalf("WithLang(en-US) = %q, want en-US or en", gotUS.Lang.String())
	}
}

func TestNilOptionsSkipped(t *testing.T) {
	var nilOpt fk.Option[NameOpts]
	got := fk.Build(nilOpt, nil, WithGender(Female), nil)
	if got.Gender != Female {
		t.Fatalf("expected %q, got %q", Female, got.Gender)
	}
	// All nil still yields defaults, no panic.
	got = fk.Build[NameOpts](nil, nil)
	if got.Gender != Male || got.Lang != language.English {
		t.Fatalf("all-nil build = %+v, want defaults", got)
	}
}

func TestWithLangNeverPanics(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("WithLang/Build panicked: %v", r)
		}
	}()
	for _, s := range []string{"", "bogus", "not-a-language", "en-US", "fa-IR", ":::;", "en"} {
		_ = fk.Build(fk.WithLang[NameOpts](s))
	}
	_ = fk.NormalizeLang("")
	_ = fk.NormalizeLang("not-a-language")
}

func TestBaseDataOverride(t *testing.T) {
	got := fk.Build(WithNickname("zed"))
	if got.Data != "zed" {
		t.Fatalf("expected Data override, got %v", got.Data)
	}
}
