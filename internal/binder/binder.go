package binder

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"

	"github.com/47monad/fakery/resources"
	"golang.org/x/text/language"
)

const (
	baseLocalesDirectory = "locales/"
)

type Binder interface {
	Bind(lang language.Tag)
}

type Data[K any] struct {
	Default  *K
	Fallback *K
}

func JSON[K any](domain string, lang language.Tag) (*Data[K], error) {
	path := func(l language.Tag) string {
		return fmt.Sprintf("%s%s/%s.json", baseLocalesDirectory, domain, l.String())
	}
	// load default
	ddata := new(K)
	if lang != language.English {
		dfd, err := fs.ReadFile(resources.LocaleFS, path(lang))
		if err != nil && !errors.Is(err, fs.ErrNotExist) {
			return nil, err
		} else if err == nil {
			if err = json.Unmarshal(dfd, &ddata); err != nil {
				return nil, err
			}
		}
	}

	var fbdata *K
	fbfd, err := fs.ReadFile(resources.LocaleFS, path(language.English))
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(fbfd, &fbdata); err != nil {
		return nil, err
	}

	if lang == language.English {
		ddata = fbdata
	}

	return &Data[K]{
		Default:  ddata,
		Fallback: fbdata,
	}, nil
}
