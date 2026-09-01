package fkrand

import (
	"errors"

	"github.com/47monad/fakery/internal/randomizer"
)

func digit() int {
	return randomizer.Digit()
}

func intWithLength(length uint32) int {
	// 10^10 - 1 does not fit in an int on 32-bit platforms, so at most
	// 9 digits can be generated portably.
	if length > 9 {
		panic(errors.New("length should be less than 10"))
	}
	if length == 0 || length == 1 {
		return randomizer.Digit()
	}
	return randomizer.Number(length)
}

func intInRange(min, max int) int {
	return randomizer.InRange(min, max)
}
