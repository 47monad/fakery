package randomizer

import (
	"math"
	"math/rand/v2"
)

func InRange(min int, max int) int {
	return getDefaultRNG().IntN(max-min) + min
}

func InRangeWith(r *rand.Rand, min, max int) int {
	return r.IntN(max-min) + min
}

func Int(max int) int {
	return InRange(0, max)
}

// Bool returns a random boolean value.
func Bool() bool {
	return getDefaultRNG().IntN(2) == 0
}

// BoolWith returns a random boolean value using the provided RNG.
func BoolWith(r *rand.Rand) bool {
	return r.IntN(2) == 0
}

// Digit returns a random digit in the inclusive range [0, 9].
func Digit() int {
	return InRange(0, 10)
}

// Number returns a random integer with exactly the given number of digits,
// i.e. a value in the inclusive range [10^(length-1), 10^length - 1].
// The maximum supported length is 9 so the result fits in an int on
// 32-bit platforms as well.
func Number(length uint32) int {
	maxLimit := int(math.Pow10(int(length))) - 1
	lowLimit := int(math.Pow10(int(length) - 1))
	return InRange(lowLimit, maxLimit+1)
}
