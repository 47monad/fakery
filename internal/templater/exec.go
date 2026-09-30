package templater

import (
	"bytes"
	"sync"
	"sync/atomic"
	"text/template"
)

// Exec is the non-panicking, parse-once replacement for MustExec.
//
// Templates come from a small closed set of locale data values and are
// parsed at most once per process. Each execution clones the cached
// template and binds the caller's funcs to the clone, so per-call state
// (e.g. a seeded *Faker) is never shared between goroutines.
//
// Unlike MustExec, Exec never panics: parse and execute failures are
// returned to the caller so domains can degrade gracefully (typically by
// falling back to the raw template string).
func Exec(name, tmpl string, funcs template.FuncMap) (string, error) {
	base, err := parseOnce(name, tmpl)
	if err == nil {
		// Cached templates are shared across calls; execute a clone with
		// the caller's function bindings so funcs are never mutated.
		exec, cerr := base.Clone()
		if cerr != nil {
			return "", cerr
		}
		if funcs != nil {
			exec = exec.Funcs(funcs)
		}
		var buf bytes.Buffer
		if eerr := exec.Execute(&buf, nil); eerr != nil {
			return "", eerr
		}
		return buf.String(), nil
	}

	// The template references a function name outside the cacheable set
	// (see placeholders). Parse it directly with the real funcs; it is
	// not cached, but it still cannot panic.
	direct, derr := template.New(name).Funcs(funcs).Parse(tmpl)
	if derr != nil {
		return "", derr
	}
	var buf bytes.Buffer
	if eerr := direct.Execute(&buf, nil); eerr != nil {
		return "", eerr
	}
	return buf.String(), nil
}

// placeholders are the template function names Exec can parse without the
// caller's funcs. Templates using only these names are cached and later
// executed on a clone with the real functions bound.
//
// New domains that introduce template function names should register a
// placeholder here so their templates are cached too; unregistered names
// still work through Exec's uncached fallback.
var placeholders = template.FuncMap{
	"prefix":    func() string { return "" },
	"suffix":    func() string { return "" },
	"firstName": func() string { return "" },
	"lastName":  func() string { return "" },
}

// cache maps a template string to its parsed representation. Keying on the
// template string (not the funcs) is what makes parsing happen once even
// though every caller supplies different function bindings.
var cache sync.Map // string -> *template.Template

// parses counts successful json-ish parses of unique template strings.
var parses atomic.Int64

func parseOnce(name, tmpl string) (*template.Template, error) {
	if v, ok := cache.Load(tmpl); ok {
		return v.(*template.Template), nil
	}
	t, err := template.New(name).Funcs(placeholders).Parse(tmpl)
	if err != nil {
		return nil, err
	}
	actual, loaded := cache.LoadOrStore(tmpl, t)
	if !loaded {
		parses.Add(1)
	}
	return actual.(*template.Template), nil
}

// ParseCount returns the number of distinct template strings parsed since
// process start (or the last ResetCache). Tests use it to prove parse-once.
func ParseCount() int64 {
	return parses.Load()
}

// ResetCache clears the template cache and the parse counter. Tests only.
func ResetCache() {
	cache.Range(func(k, _ any) bool {
		cache.Delete(k)
		return true
	})
	parses.Store(0)
}
