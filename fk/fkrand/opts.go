// Options and seed support for the fkrand domain.
//
// The random-number domain is locale-free, so Options intentionally does
// not embed fk.Base: there is no language to select. It still uses the
// shared fk.Option / fk.Build system from #11 so call sites look like
// every other domain:
//
//	f := fkrand.New(fkrand.WithSeed(42))
package fkrand

import (
	"github.com/47monad/fakery/fk"
)

// Options carries the optional seed for a Faker.
// Zero value means unseeded (non-deterministic) behavior.
type Options struct {
	Seed    uint64
	HasSeed bool
}

// WithSeed returns an Option that makes New produce a deterministic
// Faker: the same seed always yields the same sequence (for
// single-goroutine use). Without it, New seeds from crypto/rand.
func WithSeed(seed uint64) fk.Option[Options] {
	return func(o *Options) {
		if o == nil {
			return
		}
		o.Seed = seed
		o.HasSeed = true
	}
}
