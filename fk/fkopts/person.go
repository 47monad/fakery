package fkopts

import (
	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/internal/binder"
)

type NameOpts struct {
	LW     *LangWrapper
	Data   *binder.Data[fkdata.Person]
	Gender string
}

func (opts *NameOpts) Default() *NameOpts {
	opts.Gender = "male"
	return opts
}

type NameBuilder struct {
	lw   *LangWrapper
	Opts []func(*NameOpts) error
}

func Name() *NameBuilder {
	return &NameBuilder{}
}

func (b *NameBuilder) List() []func(*NameOpts) error {
	return append(b.Opts, func(op *NameOpts) error {
		op.LW = b.lw
		return nil
	})
}

func (b *NameBuilder) SetLang(lang string) *NameBuilder {
	b.lw.UseLang(lang)
	return b
}

func (b *NameBuilder) SetData(d *binder.Data[fkdata.Person]) *NameBuilder {
	b.Opts = append(b.Opts, func(op *NameOpts) error {
		op.Data = d
		return nil
	})
	return b
}

func (b *NameBuilder) SetGender(gender string) *NameBuilder {
	b.Opts = append(b.Opts, func(no *NameOpts) error {
		no.Gender = gender
		return nil
	})
	return b
}

// LastName
type LastNameOpts struct {
	LW   *LangWrapper
	Data *binder.Data[fkdata.Person]
}

type LastNameBuilder struct {
	lw   *LangWrapper
	Opts []func(*LastNameOpts) error
}

func LastName() *LastNameBuilder {
	return &LastNameBuilder{}
}

func (b *LastNameBuilder) List() []func(*LastNameOpts) error {
	return append(b.Opts, func(op *LastNameOpts) error {
		op.LW = b.lw
		return nil
	})
}

func (b *LastNameBuilder) SetLang(lang string) *LastNameBuilder {
	b.lw.UseLang(lang)
	return b
}

func (b *LastNameBuilder) SetData(d *binder.Data[fkdata.Person]) *LastNameBuilder {
	b.Opts = append(b.Opts, func(op *LastNameOpts) error {
		op.Data = d
		return nil
	})
	return b
}
