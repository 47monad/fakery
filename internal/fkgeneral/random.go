package fkgeneral

import (
	"math"
	"math/rand"
	"time"
)

func RandomElement[K any](slice []K) K {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return slice[r.Intn(len(slice))]
}

func RandomInRange(min int, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min+1) + min
}

func RandomInt(max int) int {
	return RandomInRange(0, max)
}

func RandomDigit() int {
	return RandomInRange(0, 9)
}

func RandomNumber(length uint32) int {
	maxLimit := int(math.Pow10(int(length))) - 1
	lowLimit := int(math.Pow10(int(length) - 1))
	return RandomInRange(lowLimit, maxLimit)
}
