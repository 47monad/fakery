package fktypes

type VersionRange struct {
	Min, Max int
}

type SemanticVersionOptions struct {
	Major, Minor, Patch *VersionRange
}

type SemanticVersionOption func(*SemanticVersionOptions)

func DefaultSemanticVersionOptions() *SemanticVersionOptions {
	return &SemanticVersionOptions{
		Major: &VersionRange{Min: 1, Max: 9},
		Minor: &VersionRange{Min: 1, Max: 9},
		Patch: &VersionRange{Min: 1, Max: 9},
	}
}

func WithVersionRange(part string, min, max int) SemanticVersionOption {
	return func(opts *SemanticVersionOptions) {
		switch part {
		case "major":
			opts.Major = &VersionRange{Min: min, Max: max}
		case "minor":
			opts.Minor = &VersionRange{Min: min, Max: max}
		case "patch":
			opts.Patch = &VersionRange{Min: min, Max: max}
		default:
			return
		}
	}
}
