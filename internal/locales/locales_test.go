package locales_test

import (
	"sync"
	"testing"

	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/internal/locales"
	"github.com/47monad/fakery/internal/sampler"
	"golang.org/x/text/language"
)

func TestParseOnce(t *testing.T) {
	locales.Reset()
	first := locales.Load[fkdata.Person]("person", language.English)
	if locales.ParseCount() != 1 {
		t.Fatalf("first Load parsed %d files, want 1", locales.ParseCount())
	}
	for range 10 {
		got := locales.Load[fkdata.Person]("person", language.English)
		if got != first {
			t.Fatal("expected cached pointer on repeat Load")
		}
	}
	if locales.ParseCount() != 1 {
		t.Fatalf("repeat Loads parsed %d files, want exactly 1", locales.ParseCount())
	}
}

func TestEnUSResolvesToEn(t *testing.T) {
	locales.Reset()
	en := locales.Load[fkdata.Person]("person", language.English)
	enus, err := language.Parse("en-US")
	if err != nil {
		t.Fatal(err)
	}
	got := locales.Load[fkdata.Person]("person", enus)
	if got != en {
		t.Fatal("en-US should resolve to the cached en dataset")
	}
	if len(got.Default.MaleFirstName) == 0 {
		t.Fatal("expected non-empty en person data")
	}
}

func TestUnknownReturnsEnglish(t *testing.T) {
	locales.Reset()
	en := locales.Load[fkdata.Lorem]("lorem", language.English)
	got := locales.Load[fkdata.Lorem]("lorem", language.French)
	if len(got.Default.Words) == 0 {
		t.Fatal("expected English fallback data, got empty")
	}
	if len(got.Default.Words) != len(en.Default.Words) {
		t.Fatal("unknown locale should return the English dataset")
	}
}

func TestFaMissingKeyFallsBack(t *testing.T) {
	locales.Reset()
	d := locales.Load[fkdata.Person]("person", language.Persian)
	if d == nil || d.Default == nil || d.Fallback == nil {
		t.Fatal("Load must never return nil data")
	}
	// fa.json has no "suffix" key; per-key fallback serves English.
	if len(d.Default.Suffix) != 0 {
		t.Fatalf("expected empty fa suffix, got %d entries", len(d.Default.Suffix))
	}
	if len(d.Fallback.Suffix) == 0 {
		t.Fatal("expected non-empty English fallback suffix")
	}
	pool := sampler.Select(d, func(p *fkdata.Person) []string { return p.Suffix })
	if len(pool.Suffix) == 0 {
		t.Fatal("sampler.Select should fall back to English per key")
	}
}

func TestNeverPanics(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Load panicked: %v", r)
		}
	}()
	locales.Reset()
	for _, tag := range []language.Tag{language.Und, language.English, language.Persian, language.French} {
		d := locales.Load[fkdata.Person]("person", tag)
		if d == nil || d.Default == nil || d.Fallback == nil {
			t.Fatalf("Load(%q) returned nil", tag.String())
		}
	}
	// Unknown domain degrades to empty non-nil data.
	d := locales.Load[fkdata.Person]("nonexistent", language.English)
	if d == nil || d.Default == nil || d.Fallback == nil {
		t.Fatal("unknown domain must return non-nil empty data")
	}
}

func TestConcurrentLoad(t *testing.T) {
	locales.Reset()
	var wg sync.WaitGroup
	for range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			d := locales.Load[fkdata.Person]("person", language.Persian)
			if d == nil || d.Default == nil || d.Fallback == nil {
				t.Error("concurrent Load returned nil")
			}
		}()
	}
	wg.Wait()
}
