package locales_test

import (
	"testing"

	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/internal/binder"
	"github.com/47monad/fakery/internal/locales"
	"golang.org/x/text/language"
)

// Old path: file read + JSON parse on every call.
func BenchmarkBinderJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d, err := binder.JSON[fkdata.Person]("person", language.Persian)
		if err != nil {
			b.Fatal(err)
		}
		_ = d
	}
}

// New path: parse once, then cached pointer fetch.
func BenchmarkLoadCached(b *testing.B) {
	_ = locales.Load[fkdata.Person]("person", language.Persian) // warmup
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = locales.Load[fkdata.Person]("person", language.Persian)
	}
}
