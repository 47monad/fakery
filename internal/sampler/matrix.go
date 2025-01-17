package sampler

import (
	"strings"

	"github.com/47monad/fakery/internal/randomizer"
)

func makeMatrixPool[K any](d [][]K) [][]K {
	var res [][]K
	for j := 0; j < 5; j++ {
		ph := make([]K, len(d))
		for i, s := range d {
			ph[i] = randomizer.Element(s)
		}
		res = append(res, ph)
	}
	return res
}

func KeyMatrix(d [][]string) []string {
	_d := makeMatrixPool(d)
	if _d == nil {
		return []string{}
	}
	var res []string
	for _, row := range _d {
		if r := strings.Join(row, " "); r != "" {
			res = append(res, r)
		}
	}
	return res
}
