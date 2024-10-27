package fklocaler

import (
	"github.com/47monad/fakery/internal/fkgeneral"
)

func Sample[K any, S any](d *K, keyer func(*K) []S) S {
	return fkgeneral.RandomElement(keyer(d))
}
