package fktemplater

import (
	"bytes"
	"text/template"
)

func MustExecTemplate(name string, s string, gfmap template.FuncMap) string {
	tmpl, err := template.New(name).Funcs(gfmap).Parse(s)
	if err != nil {
		panic(err)
	}
	var res bytes.Buffer
	err = tmpl.Execute(&res, nil)
	if err != nil {
		panic(err)
	}
	return res.String()
}
