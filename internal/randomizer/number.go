package randomizer

import (
	"math"
	"math/rand/v2"
	"time"
)

func InRange(min int, max int) int {
	r := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 17))
	return r.IntN(max-min+1) + min
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
