package person_test

import (
	"sync"
	"testing"

	"github.com/47monad/fakery/internal/locales"
	"github.com/47monad/fakery/person"
	"golang.org/x/text/language"
)

// --- helpers ---------------------------------------------------------------

func contains(list []string, v string) bool {
	for _, s := range list {
		if s == v {
			return true
		}
	}
	return false
}

func hasPersian(s string) bool {
	for _, r := range s {
		// Arabic/Persian block; the locale data lives here.
		if r >= '\u0600' && r <= '\u06FF' {
			return true
		}
	}
	return false
}

// --- target API ------------------------------------------------------------

// TestTargetAPI mirrors the exact three-tier API from the issue.
func TestTargetAPI(t *testing.T) {
	p := person.New(person.WithLang("fa"), person.WithSeed(42))
	if got := p.Name(person.WithGender(person.Female)); got == "" {
		t.Fatal("tier-2 Name(Female) returned empty")
	}

	if got := person.FirstName(); got == "" {
		t.Fatal("tier-3 FirstName returned empty")
	}
	if got := person.LastName(person.WithLang("fa")); got == "" {
		t.Fatal("tier-3 LastName(fa) returned empty")
	}
}

// --- determinism -----------------------------------------------------------

func TestSeededDeterminism(t *testing.T) {
	drain := func(f *person.Faker) []string {
		out := make([]string, 0, 32)
		for range 4 {
			out = append(out,
				f.FirstName(), f.LastName(), f.Name(),
				f.NameWithMiddle(), f.Prefix(), f.Suffix(),
			)
		}
		return out
	}
	a := drain(person.New(person.WithSeed(7)))
	b := drain(person.New(person.WithSeed(7)))
	if len(a) != len(b) {
		t.Fatalf("length mismatch: %d vs %d", len(a), len(b))
	}
	for i := range a {
		if a[i] != b[i] {
			t.Fatalf("same seed diverged at %d: %q vs %q", i, a[i], b[i])
		}
	}
	c := drain(person.New(person.WithSeed(8)))
	same := true
	for i := range a {
		if a[i] != c[i] {
			same = false
			break
		}
	}
	if same {
		t.Fatal("different seeds produced identical sequences")
	}
}

// TestGoldenOutput pins the seeded stream to fixed expected strings. This is
// the cross-run determinism guarantee: re-running the suite must produce the
// same values.
func TestGoldenOutput(t *testing.T) {
	en := []string{
		"Ms. Donald Wunsch",
		"Zack Johnston",
		"Amb. Kena Conn",
		"Sallie O'Hara",
		"Fred Bergstrom",
		"Prof. Tianna Windler",
		"Darwin Ebert",
		"Peggie Witting",
	}
	f := person.New(person.WithSeed(42))
	for i, want := range en {
		if got := f.Name(); got != want {
			t.Fatalf("en Name #%d = %q, want %q", i, got, want)
		}
	}

	fa := []string{
		"آقای مشکان\u200cدخت قانونی یادگار",
		"بوسه یلدا",
		"آبگینه هومن",
		"سمانه غنی",
	}
	g := person.New(person.WithSeed(42), person.WithLang("fa"))
	for i, want := range fa {
		if got := g.Name(person.WithGender(person.Female)); got != want {
			t.Fatalf("fa Name #%d = %q, want %q", i, got, want)
		}
	}
}

// --- locale + per-key fallback --------------------------------------------

func TestPersianLocale(t *testing.T) {
	f := person.New(person.WithLang("fa"), person.WithSeed(1))
	for range 50 {
		if got := f.LastName(); !hasPersian(got) {
			t.Fatalf("LastName(fa) = %q, want Persian characters", got)
		}
		if got := f.FirstName(person.WithGender(person.Female)); !hasPersian(got) {
			t.Fatalf("FirstName(fa, Female) = %q, want Persian characters", got)
		}
	}
}

func TestPerKeyFallbackToEnglish(t *testing.T) {
	en := locales.Load[person.Locale]("person", language.English).Default
	fa := locales.Load[person.Locale]("person", language.Persian).Default

	// Sanity: fa really does omit suffix/nameWithMiddle.
	if len(fa.Suffix) != 0 {
		t.Fatalf("expected fa to omit suffix, got %d entries", len(fa.Suffix))
	}
	if len(fa.NameWithMiddle) != 0 {
		t.Fatalf("expected fa to omit nameWithMiddle, got %d entries", len(fa.NameWithMiddle))
	}

	// Missing keys fall back per key to English.
	f := person.New(person.WithLang("fa"), person.WithSeed(3))
	for range 50 {
		if got := f.Suffix(); !contains(en.Suffix, got) {
			t.Fatalf("Suffix(fa) = %q, want an English suffix", got)
		}
		if got := f.NameWithMiddle(); got == "" {
			t.Fatal("NameWithMiddle(fa) fell back to an empty string")
		}
		if got := f.Prefix(); !contains(fa.Prefix, got) {
			t.Fatalf("Prefix(fa) = %q, want a fa prefix (present key must win)", got)
		}
	}
}

// --- gender ----------------------------------------------------------------

func TestGenderAnyUsesBothPools(t *testing.T) {
	en := locales.Load[person.Locale]("person", language.English).Default
	f := person.New(person.WithSeed(11)) // default GenderAny
	seenMale, seenFemale := false, false
	for range 500 {
		got := f.FirstName()
		if contains(en.MaleFirstName, got) {
			seenMale = true
		}
		if contains(en.FemaleFirstName, got) {
			seenFemale = true
		}
	}
	if !seenMale || !seenFemale {
		t.Fatalf("GenderAny pools: male=%v female=%v, want both", seenMale, seenFemale)
	}
}

func TestFemaleNeverYieldsMaleOnlyNames(t *testing.T) {
	en := locales.Load[person.Locale]("person", language.English).Default
	f := person.New(person.WithSeed(12), person.WithGender(person.Female))
	for range 500 {
		got := f.FirstName()
		if !contains(en.FemaleFirstName, got) {
			t.Fatalf("Female FirstName = %q, not in the female pool", got)
		}
	}
}

func TestMaleNeverYieldsFemaleOnlyNames(t *testing.T) {
	en := locales.Load[person.Locale]("person", language.English).Default
	f := person.New(person.WithSeed(13), person.WithGender(person.Male))
	for range 500 {
		got := f.FirstName()
		if !contains(en.MaleFirstName, got) {
			t.Fatalf("Male FirstName = %q, not in the male pool", got)
		}
	}
}

func TestGenderNormalization(t *testing.T) {
	if person.ParseGender("MALE") != person.Male {
		t.Fatal("ParseGender(MALE) != Male")
	}
	if person.ParseGender("FEMALE") != person.Female {
		t.Fatal("ParseGender(FEMALE) != Female")
	}
	if person.ParseGender("typo") != person.GenderAny {
		t.Fatal("ParseGender(typo) should be GenderAny, not male")
	}
	if person.WithGender(person.Gender(99)) == nil {
		t.Fatal("WithGender returned nil")
	}
	// An out-of-range gender must normalize and never panic.
	f := person.New(person.WithGender(person.Gender(99)), person.WithSeed(1))
	if got := f.FirstName(); got == "" {
		t.Fatal("normalized gender produced empty name")
	}
	if person.Female.String() != "female" || person.Male.String() != "male" || person.GenderAny.String() != "any" {
		t.Fatal("Gender.String mismatch")
	}
	if !person.Female.IsValid() || person.Gender(42).IsValid() {
		t.Fatal("Gender.IsValid mismatch")
	}
}

// --- custom data -----------------------------------------------------------

func TestCustomDataOverride(t *testing.T) {
	custom := person.Locale{
		MaleFirstName: []string{"Zxcvbn"},
		LastName:      []string{"Qwerty"},
	}
	if got := person.FirstName(person.WithData(custom), person.WithGender(person.Male)); got != "Zxcvbn" {
		t.Fatalf("WithData FirstName = %q, want %q", got, "Zxcvbn")
	}
	if got := person.LastName(person.WithPersonData(&custom)); got != "Qwerty" {
		t.Fatalf("WithPersonData LastName = %q, want %q", got, "Qwerty")
	}
	if got := person.Name(person.WithData(custom)); got == "" {
		t.Fatal("WithData Name returned empty")
	}
}

func TestInvalidCustomDataFallsBack(t *testing.T) {
	// Invalid type, empty locale, and nil must all fall back to the embedded
	// data without panicking.
	for _, data := range []any{42, "nope", struct{}{}, person.Locale{}, (*person.Locale)(nil), []int{1}} {
		if got := person.Name(person.WithLang("fa"), person.WithData(data)); got == "" {
			t.Fatalf("WithData(%v) produced empty Name", data)
		}
	}
}

// --- robustness ------------------------------------------------------------

func TestNoPanicsOnOptionCombinations(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("panicked on option combination: %v", r)
		}
	}()

	langs := []string{"", "en", "fa", "fa-IR", "en-US", "not-a-language", ":::;"}
	genders := []person.Gender{person.GenderAny, person.Male, person.Female, person.Gender(-7)}
	datas := []any{nil, 1, person.Locale{}, person.Locale{MaleFirstName: []string{"X"}}}

	for _, lang := range langs {
		for _, g := range genders {
			for _, d := range datas {
				f := person.New(
					person.WithLang(lang),
					person.WithGender(g),
					person.WithData(d),
					person.WithSeed(1),
				)
				if f.Name() == "" && d == nil {
					t.Fatalf("Name empty for lang=%q gender=%v data=%v", lang, g, d)
				}
				_ = f.FirstName()
				_ = f.LastName()
				_ = f.NameWithMiddle()
				_ = f.Prefix()
				_ = f.Suffix()
			}
		}
	}

	// Nil options are skipped.
	_ = person.New(nil, nil, person.WithLang("fa"), nil).Name()

	// Nil receiver methods must not panic.
	var nilFaker *person.Faker
	_ = nilFaker.FirstName()
	_ = nilFaker.LastName()
	_ = nilFaker.Name()
	_ = nilFaker.NameWithMiddle()
	_ = nilFaker.Prefix()
	_ = nilFaker.Suffix()
}

func TestTier3Smoke(t *testing.T) {
	for range 100 {
		if person.FirstName() == "" {
			t.Fatal("tier-3 FirstName empty")
		}
		if person.LastName() == "" {
			t.Fatal("tier-3 LastName empty")
		}
		if person.Name() == "" {
			t.Fatal("tier-3 Name empty")
		}
		if person.NameWithMiddle() == "" {
			t.Fatal("tier-3 NameWithMiddle empty")
		}
		if person.Prefix() == "" {
			t.Fatal("tier-3 Prefix empty")
		}
		if person.Suffix() == "" {
			t.Fatal("tier-3 Suffix empty")
		}
	}
}

func TestConcurrentUse(t *testing.T) {
	f := person.New(person.WithSeed(1), person.WithLang("fa"))
	var wg sync.WaitGroup
	for range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 500 {
				_ = f.FirstName(person.WithGender(person.Female))
				_ = f.LastName()
				_ = f.Name()
				_ = f.NameWithMiddle()
				_ = f.Prefix()
				_ = f.Suffix()
			}
		}()
	}
	wg.Wait()
}
