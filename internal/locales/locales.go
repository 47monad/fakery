// Package locales is the parse-once locale cache for all faker domains.
//
// Every standalone call used to hit the embedded filesystem twice
// (target locale + English fallback) with fs.ReadFile + json.Unmarshal.
// This package reads and unmarshals each (domain, lang) file exactly
// once per process and serves cached pointers afterwards.
//
// Language matching replaces exact filename lookup so regional tags
// resolve sanely: en-US -> en, fa-IR -> fa. Unknown or malformed tags
// fall back to English. Load never panics and never returns an error
// or nil data — English is always the answer of last resort.
package locales

import (
	"encoding/json"
	"io/fs"
	"path"
	"sort"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/47monad/fakery/internal/binder"
	"github.com/47monad/fakery/resources"
	"golang.org/x/text/language"
)

// cache stores composed datasets: key "domain|resolved" -> any (*binder.Data[K]).
// The English entry ("domain|en") is populated on the first miss so the
// fallback file is never parsed twice.
var cache sync.Map

// matchers stores per-domain matchers: domain -> *domainMatcher.
var matchers sync.Map

// mu serializes cache misses so each file is parsed exactly once
// even under concurrent first use.
var mu sync.Mutex

// parses counts json.Unmarshal executions (one per file read hit).
// Exposed via ParseCount for tests.
var parses atomic.Int64

type domainMatcher struct {
	tags    []language.Tag
	keys    []string // filename stems, parallel to tags
	matcher language.Matcher
}

// Reset clears all caches and the parse counter. For tests only.
func Reset() {
	clearMap(&cache)
	clearMap(&matchers)
	parses.Store(0)
}

func clearMap(m *sync.Map) {
	m.Range(func(k, _ any) bool {
		m.Delete(k)
		return true
	})
}

// ParseCount returns the number of JSON parses performed since
// process start (or the last Reset).
func ParseCount() int64 {
	return parses.Load()
}

// Load returns the dataset for domain in the requested language.
// Default holds the resolved locale, Fallback the English data so
// per-key fallback in sampler.Select keeps working. If the resolved
// locale is English, both point to the same value.
//
// It never panics, never returns an error, and never returns nil:
// missing files or malformed JSON degrade to (empty) English data.
func Load[K any](domain string, tag language.Tag) *binder.Data[K] {
	key := resolve(domain, tag)
	cacheKey := domain + "|" + key

	if v, ok := cache.Load(cacheKey); ok {
		if d, ok := v.(*binder.Data[K]); ok && d != nil {
			return d
		}
		// Type mismatch: same domain requested with a different K.
		// Fall through and reload (last writer wins). This only
		// happens on programmer error; domains map 1:1 to types.
	}

	mu.Lock()
	defer mu.Unlock()

	// Double-check under the lock.
	if v, ok := cache.Load(cacheKey); ok {
		if d, ok := v.(*binder.Data[K]); ok && d != nil {
			return d
		}
	}

	en := englishData[K](domain)
	var d *binder.Data[K]
	if key == "en" {
		d = &binder.Data[K]{Default: en, Fallback: en}
	} else {
		d = &binder.Data[K]{Default: loadFile[K](domain, key), Fallback: en}
	}
	cache.Store(cacheKey, d)
	return d
}

// englishData returns the cached English payload for domain,
// parsing and caching it on first use. Must be called with mu held.
func englishData[K any](domain string) *K {
	const suffix = "|en"
	if v, ok := cache.Load(domain + suffix); ok {
		if d, ok := v.(*binder.Data[K]); ok && d != nil && d.Fallback != nil {
			return d.Fallback
		}
	}
	en := loadFile[K](domain, "en")
	cache.Store(domain+suffix, &binder.Data[K]{Default: en, Fallback: en})
	return en
}

// loadFile reads and unmarshals one locale file. Missing files and
// malformed JSON yield an empty value — never an error, never nil.
func loadFile[K any](domain, key string) *K {
	out := new(K)
	raw, err := fs.ReadFile(resources.LocaleFS, "locales/"+domain+"/"+key+".json")
	if err != nil {
		return out
	}
	parses.Add(1)
	if err := json.Unmarshal(raw, out); err != nil {
		return new(K)
	}
	return out
}

// resolve maps a requested tag to a filename stem for domain.
// Unknown, undetermined, or unmatched tags yield "en".
func resolve(domain string, tag language.Tag) string {
	if tag == language.Und || tag.IsRoot() {
		return "en"
	}
	m := matcherFor(domain)
	_, index, conf := m.matcher.Match(tag)
	if conf == language.No {
		return "en"
	}
	if index < 0 || index >= len(m.keys) {
		return "en"
	}
	return m.keys[index]
}

// matcherFor builds (once per domain) a matcher over the locale
// files actually embedded for that domain.
func matcherFor(domain string) *domainMatcher {
	if v, ok := matchers.Load(domain); ok {
		if m, ok := v.(*domainMatcher); ok && m != nil {
			return m
		}
	}

	mu.Lock()
	defer mu.Unlock()

	if v, ok := matchers.Load(domain); ok {
		if m, ok := v.(*domainMatcher); ok && m != nil {
			return m
		}
	}

	m := buildMatcher(domain)
	matchers.Store(domain, m)
	return m
}

func buildMatcher(domain string) *domainMatcher {
	files, err := fs.Glob(resources.LocaleFS, "locales/"+domain+"/*.json")
	keys := make([]string, 0, 4)
	seen := make(map[string]bool, 4)
	if err == nil {
		for _, f := range files {
			stem := strings.TrimSuffix(path.Base(f), ".json")
			if stem == "" || seen[stem] {
				continue
			}
			// Only advertise stems that parse as language tags.
			if _, err := language.Parse(stem); err != nil {
				continue
			}
			seen[stem] = true
			keys = append(keys, stem)
		}
	}
	// English is always the fallback; guarantee its presence even if
	// the domain ships no files (Load then serves empty data).
	if !seen["en"] {
		keys = append(keys, "en")
	}
	sort.Strings(keys)

	tags := make([]language.Tag, len(keys))
	for i, k := range keys {
		tags[i], _ = language.Parse(k)
	}
	return &domainMatcher{tags: tags, keys: keys, matcher: language.NewMatcher(tags)}
}
