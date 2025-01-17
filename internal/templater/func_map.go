package templater

import (
	"html/template"
)

func MergeFuncMaps(mapList ...template.FuncMap) template.FuncMap {
	res := template.FuncMap{}
	for _, m := range mapList {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}
