package fkrand

import (
	"errors"

	"github.com/47monad/fakery/internal/randomizer"
)

func digit() int {
	return randomizer.Digit()
}

func intWithLength(length uint32) int {
	if length > 10 {
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
