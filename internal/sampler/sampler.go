package sampler

import (
	"github.com/47monad/fakery/internal/binder"
	"github.com/47monad/fakery/internal/randomizer"
)

func Select[K any, S any](d *binder.Data[K], keyer func(*K) []S) *K {
	pool := keyer(d.Default)
	data := d.Default
	if len(pool) == 0 {
		pool = keyer(d.Fallback)
		data = d.Fallback
	}
	return data
}

func Run[K any, S any](d *binder.Data[K], keyer func(*K) []S) S {
	pool := Select(d, keyer)
	return randomizer.Element(keyer(pool))
}

func RunMany[K any, S any](count int, d *binder.Data[K], keyer func(*K) []S) []S {
	pool := Select(d, keyer)
	return randomizer.Elements(count, keyer(pool))
}
