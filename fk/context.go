package fk

import (
	"github.com/47monad/fakery/internal/binder"
	"golang.org/x/text/language"
)

type Context[K any] struct {
	Lang   language.Tag
	Data   *binder.Data[K]
	binder func(language.Tag) *binder.Data[K]
}

func (c *Context[K]) SetBinder(fn func(language.Tag) *binder.Data[K]) *Context[K] {
	c.binder = fn
	c.Data = c.binder(c.Lang)
	return c
}

func DefaultContext[K any]() *Context[K] {
	return &Context[K]{
		Lang: language.English,
		Data: nil,
	}
}

func NewContext[K any](lang language.Tag, data *binder.Data[K]) *Context[K] {
	b := DefaultContext[K]()
	if lang != language.Und {
		b.Lang = lang
	}
	return b
}
