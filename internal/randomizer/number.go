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

func Digit() int {
	return InRange(0, 9)
}

func Number(length uint32) int {
	maxLimit := int(math.Pow10(int(length))) - 1
	lowLimit := int(math.Pow10(int(length) - 1))
	return InRange(lowLimit, maxLimit)
}
