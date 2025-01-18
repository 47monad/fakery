package fkopts

import "github.com/47monad/fakery/internal/binder"

type GeneralOpts[K any] struct {
	LW   *LangWrapper
	Data *binder.Data[K]
}

type GeneralBuilder[K any] struct {
	lw   *LangWrapper
	Opts []func(*GeneralOpts[K]) error
}

func General[K any]() *GeneralBuilder[K] {
	return &GeneralBuilder[K]{lw: &LangWrapper{}}
}

func (b *GeneralBuilder[K]) List() []func(*GeneralOpts[K]) error {
	return append(b.Opts, func(op *GeneralOpts[K]) error {
		op.LW = b.lw
		return nil
	})
}

func (b *GeneralBuilder[K]) SetLang(lang string) *GeneralBuilder[K] {
	b.lw.UseLang(lang)
	return b
}

func (b *GeneralBuilder[K]) SetData(d *binder.Data[K]) *GeneralBuilder[K] {
	b.Opts = append(b.Opts, func(op *GeneralOpts[K]) error {
		op.Data = d
		return nil
	})
	return b
}
