// Seeded Faker for the fkrand domain (issue #17).
//
// Faker holds an *internal/rng.RNG so callers get reproducible streams
// via fkrand.New(fkrand.WithSeed(42)). Package-level (tier-3) functions
// in solo.go delegate to a default unseeded Faker for zero-setup use.
package fkrand

import (
	"errors"

	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/internal/rng"
)

// Faker is a seeded source of random numbers and random picks.
// The zero value is not usable; construct with New.
// It is safe for concurrent use (the underlying RNG holds a mutex),
// but only single-goroutine use is deterministic in ordering.
type Faker struct {
	r *rng.RNG
}

// New returns a Faker. With no options it is seeded non-deterministically
// (crypto/rand, falling back to time); with WithSeed it is deterministic:
//
//	fkrand.New(fkrand.WithSeed(42)) // same seed → same sequence
func New(opts ...fk.Option[Options]) *Faker {
	o := fk.Build(opts...)
	if o.HasSeed {
		return &Faker{r: rng.New(o.Seed)}
	}
	return &Faker{r: rng.Default()}
}

// rngOrDefault guards against a nil Faker or nil RNG so methods never
// panic on a nil receiver; callers get default (unseeded) behavior.
func (f *Faker) rngOrDefault() *rng.RNG {
	if f != nil && f.r != nil {
		return f.r
	}
	return rng.Default()
}

// Digit returns a random digit in the inclusive range [0, 9].
func (f *Faker) Digit() int {
	return f.rngOrDefault().IntN(10)
}

// pow10 returns 10^n for n in [0, 10] using integer arithmetic.
func pow10(n uint32) int {
	p := 1
	for range n {
		p *= 10
	}
	return p
}

// Int returns a random integer with exactly the given number of digits:
// length 0 or 1 yields a single digit in [0, 9]; length 2 yields
// [10, 99], and so on up to length 9 ([100000000, 999999999]).
// A length greater than 9 panics: 10^10-1 does not fit in an int on
// 32-bit platforms, so at most 9 digits can be generated portably.
func (f *Faker) Int(length uint32) int {
	if length > 9 {
		panic(errors.New("length should be less than 10"))
	}
	if length == 0 || length == 1 {
		return f.Digit()
	}
	low := pow10(length - 1)
	high := pow10(length) - 1
	return f.IntInRange(low, high+1)
}

// IntInRange returns a random integer in the half-open range [min, max).
// It NEVER returns max. If max <= min it returns min (no panic), so
// IntInRange(a, a) == a and reversed bounds collapse to min.
func (f *Faker) IntInRange(min, max int) int {
	if max <= min {
		return min
	}
	// int64 arithmetic so the full 32-bit int range never overflows
	// (only the single 64-bit extreme span of 2^64-1 wraps, which no
	// caller can meaningfully sample anyway).
	return int(f.rngOrDefault().Int64N(int64(max)-int64(min))) + min
}

// Float64 returns a uniform value in [0.0, 1.0).
func (f *Faker) Float64() float64 {
	return f.rngOrDefault().Float64()
}

// Float64InRange returns a uniform value in [min, max).
// If max <= min it returns min (no panic), mirroring IntInRange.
func (f *Faker) Float64InRange(min, max float64) float64 {
	if !(max > min) {
		return min
	}
	return min + f.rngOrDefault().Float64()*(max-min)
}

// Bool returns a random boolean value.
func (f *Faker) Bool() bool {
	return f.rngOrDefault().Bool()
}

// fakerOf resolves the variadic Faker argument shared by Element and
// Elements: the first non-nil entry wins, otherwise the default.
func fakerOf(fakers ...*Faker) *Faker {
	for _, f := range fakers {
		if f != nil && f.r != nil {
			return f
		}
	}
	return defaultFaker
}

// defaultFaker backs all tier-3 package-level functions. It uses an
// unseeded RNG, so tier-3 output is non-deterministic by design.
var defaultFaker = New()
