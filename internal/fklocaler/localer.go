package fklocaler

import (
	"github.com/47monad/fakery/internal/fktypes"
)

func Localize[K any, S any](d *fktypes.ResultData[K], keyer func(*K) []S) (*K, func() S) {
	pool := keyer(d.Default)
	data := d.Default
	if len(pool) == 0 {
		pool = keyer(d.Fallback)
		data = d.Fallback
	}
	return data, func() S { return Sample(data, keyer) }
}
