package fknumber

import "github.com/47monad/fakery/internal/fkgeneral"

func NumTFunc(args ...interface{}) (interface{}, error) {
	length := uint32(1)
	if args[0] != nil {
		length = uint32(args[0].(int))
	}
	return fkgeneral.RandomNumber(length), nil
}
