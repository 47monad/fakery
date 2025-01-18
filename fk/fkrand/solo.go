package fkrand

import "github.com/47monad/fakery/internal/randomizer"

func Digit() int {
	return digit()
}

func Int(length uint32) int {
	return intWithLength(length)
}

func IntInRange(min, max int) int {
	return intInRange(min, max)
}

func Bool() bool {
	return intInRange(0, 2) == 1
}

func Element[K any](slice []K) K {
	return randomizer.Element(slice)
}

func Elements[K any](count int, slice []K) []K {
	return randomizer.Elements(count, slice)
}
