package fknumber

import (
	"errors"

	"github.com/47monad/fakery/internal/fkgeneral"
	"github.com/47monad/fakery/internal/fktypes"
)

type NumberGenerator struct {
	Digit  func() int
	Number func(uint32) int
}

func NewNumberGenerator(opts *fktypes.FakeryConfig) NumberGenerator {
	return NumberGenerator{
		Digit:  generateDigit(opts),
		Number: generateNumber(opts),
	}
}

func generateDigit(opts *fktypes.FakeryConfig) func() int {
	return func() int {
		return fkgeneral.RandomDigit()
	}
}

func generateNumber(opts *fktypes.FakeryConfig) func(uint32) int {
	return func(length uint32) int {
		if length > 10 {
			panic(errors.New("length should be less than 10"))
		}
		if length == 0 || length == 1 {
			return generateDigit(opts)()
		}
		return fkgeneral.RandomNumber(length)
	}
}
