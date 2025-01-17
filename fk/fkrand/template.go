package fkrand

import "text/template"

func TemplateFuncMap() template.FuncMap {
	return template.FuncMap{
		"num": numTFunc,
	}
}

func numTFunc(args ...interface{}) (interface{}, error) {
	length := uint32(1)
	if args[0] != nil {
		length = uint32(args[0].(int))
	}
	return intWithLength(length), nil
}
