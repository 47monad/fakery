package fkopts

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
