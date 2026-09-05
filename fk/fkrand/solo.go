// Tier-3 package-level generators backed by the default unseeded Faker.
//
// For deterministic output, construct a Faker with WithSeed and either
// call its methods or pass it as the trailing argument to Element and
// Elements:
//
//	f := fkrand.New(fkrand.WithSeed(42))
//	f.Digit()                  // deterministic
//	fkrand.Element(colors, f)  // deterministic pick
package fkrand

// Digit returns a random digit in the inclusive range [0, 9]
// using the default RNG.
func Digit() int {
	return defaultFaker.Digit()
}

// Int returns a random integer with exactly the given number of digits.
// A length of 0 or 1 yields a single digit. The maximum supported
// length is 9; larger values panic.
func Int(length uint32) int {
	return defaultFaker.Int(length)
}

// IntInRange returns a random integer in the half-open range [min, max).
// It NEVER returns max; max <= min yields min.
func IntInRange(min, max int) int {
	return defaultFaker.IntInRange(min, max)
}

// Float64 returns a uniform value in [0.0, 1.0) using the default RNG.
func Float64() float64 {
	return defaultFaker.Float64()
}

// Float64InRange returns a uniform value in [min, max) using the
// default RNG. max <= min yields min.
func Float64InRange(min, max float64) float64 {
	return defaultFaker.Float64InRange(min, max)
}

// Bool returns a random boolean value using the default RNG.
func Bool() bool {
	return defaultFaker.Bool()
}

// Element returns a random element of slice using the default RNG.
// If slice is empty it returns the zero value of K instead of
// panicking. Pass an optional *Faker for the seeded path; a nil Faker
// is ignored.
//
//	fkrand.Element(colors)    // tier 3: default RNG
//	fkrand.Element(colors, f) // tier 2: seeded via f
func Element[K any](slice []K, fakers ...*Faker) K {
	if len(slice) == 0 {
		var zero K
		return zero
	}
	return slice[fakerOf(fakers...).rngOrDefault().IntN(len(slice))]
}

// Elements returns count distinct elements picked at random from slice
// using the default RNG. If count > len(slice) it is clamped to
// len(slice); if count <= 0 or slice is empty it returns an empty
// (non-nil) slice. It never panics. Pass an optional *Faker for the
// seeded path; a nil Faker is ignored.
func Elements[K any](count int, slice []K, fakers ...*Faker) []K {
	if count > len(slice) {
		count = len(slice)
	}
	if count < 0 {
		count = 0
	}
	f := fakerOf(fakers...)
	result := make([]K, 0, count)
	used := make(map[int]bool, count)
	for len(result) < count {
		i := f.rngOrDefault().IntN(len(slice))
		if !used[i] {
			result = append(result, slice[i])
			used[i] = true
		}
	}
	return result
}
