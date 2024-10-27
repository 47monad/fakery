package fkloader

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"

	"github.com/47monad/fakery/internal/fktypes"
	"golang.org/x/text/language"
)

const (
	baseLocalesDirectory = "resources/locales/"
)

func FromJSON[K any](fsys fs.FS, domain string, lang language.Tag) (*fktypes.ResultData[K], error) {
	path := func(l language.Tag) string {
		return fmt.Sprintf("%s%s/%s.json", baseLocalesDirectory, domain, l.String())
	}
	// load default
	ddata := new(K)
	if lang != language.English {
		dfd, err := fs.ReadFile(fsys, path(lang))
		if err != nil && !errors.Is(err, fs.ErrNotExist) {
			return nil, err
		} else if err == nil {
			if err = json.Unmarshal(dfd, &ddata); err != nil {
				return nil, err
			}
		}
	}

	var fbdata *K
	fbfd, err := fs.ReadFile(fsys, path(language.English))
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(fbfd, &fbdata); err != nil {
		return nil, err
	}

	return &fktypes.ResultData[K]{
		Default:  ddata,
		Fallback: fbdata,
	}, nil
}
