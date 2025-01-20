package fkopts

import (
	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/internal/binder"
)

type VersionRange struct {
	Min, Max int
}

type SemVerOpts struct {
	Major, Minor, Patch *VersionRange
}

func (opts *SemVerOpts) Default() *SemVerOpts {
	opts.Major = &VersionRange{Min: 1, Max: 9}
	opts.Minor = &VersionRange{Min: 1, Max: 9}
	opts.Patch = &VersionRange{Min: 1, Max: 9}
	return opts
}

type SemVerBuilder struct {
	Opts []func(*SemVerOpts) error
}

func SemVer() *SemVerBuilder {
	return &SemVerBuilder{}
}

func (b *SemVerBuilder) List() []func(*SemVerOpts) error {
	return b.Opts
}

func (b *SemVerBuilder) SetMajorRange(min, max int) *SemVerBuilder {
	b.Opts = append(b.Opts, func(svo *SemVerOpts) error {
		svo.Major = &VersionRange{Min: min, Max: max}

		return nil
	})

	return b
}

func (b *SemVerBuilder) SetMinorRange(min, max int) *SemVerBuilder {
	b.Opts = append(b.Opts, func(svo *SemVerOpts) error {
		svo.Minor = &VersionRange{Min: min, Max: max}

		return nil
	})

	return b
}

func (b *SemVerBuilder) SetPatchRange(min, max int) *SemVerBuilder {
	b.Opts = append(b.Opts, func(svo *SemVerOpts) error {
		svo.Patch = &VersionRange{Min: min, Max: max}

		return nil
	})

	return b
}

// App Name Options
type AppNameOpts struct {
	LW   *LangWrapper
	Data *binder.Data[fkdata.App]
}

type AppNameBuilder struct {
	lw   *LangWrapper
	Opts []func(*AppNameOpts) error
}

func AppName() *AppNameBuilder {
	return &AppNameBuilder{lw: &LangWrapper{}}
}

func (b *AppNameBuilder) List() []func(*AppNameOpts) error {
	return append(b.Opts, func(op *AppNameOpts) error {
		op.LW = b.lw
		return nil
	})
}

func (b *AppNameBuilder) SetLang(lang string) *AppNameBuilder {
	b.lw.UseLang(lang)
	return b
}

func (b *AppNameBuilder) SetData(data *binder.Data[fkdata.App]) *AppNameBuilder {
	b.Opts = append(b.Opts, func(op *AppNameOpts) error {
		op.Data = data
		return nil
	})
	return b
}
