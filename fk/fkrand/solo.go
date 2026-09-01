package fkrand

import "github.com/47monad/fakery/internal/randomizer"

// Digit returns a random digit in the inclusive range [0, 9].
func Digit() int {
	return digit()
}

// Int returns a random integer with exactly the given number of digits.
// A length of 0 or 1 yields a single digit. The maximum supported
// length is 9; larger values panic.
func Int(length uint32) int {
	return intWithLength(length)
}

// IntInRange returns a random integer in the half-open range [min, max).
func IntInRange(min, max int) int {
	return intInRange(min, max)
}

// Bool returns a random boolean value.
func Bool() bool {
	return randomizer.Bool()
}

func Element[K any](slice []K) K {
	return randomizer.Element(slice)
}

func Elements[K any](count int, slice []K) []K {
	return randomizer.Elements(count, slice)
}
