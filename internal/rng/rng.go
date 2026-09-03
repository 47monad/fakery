// Package rng provides a seeded, thread-safe random number generator
// for fakery.
//
// A seeded RNG used from multiple goroutines is safe (all methods hold
// an internal mutex) but nondeterministic in ordering: concurrent callers
// race for the mutex, so the interleaving — and therefore the sequence
// observed by any single goroutine — varies run to run.
//
// Determinism is guaranteed for single-goroutine use: two RNGs created
// with the same seed produce identical sequences.
//
// Per-domain streams (see Derived) give each faker domain an independent,
// call-order-stable sequence derived from one master seed.
package rng

import (
	crand "crypto/rand"
	"encoding/binary"
	"hash/fnv"
	"math/rand/v2"
	"sync"
	"time"
)

// pcgStream is the fixed PCG stream ID used for all seeded RNGs.
// Keeping it constant means the seed alone determines the sequence.
const pcgStream uint64 = 17

// RNG is a mutex-guarded wrapper around math/rand/v2.Rand.
// math/rand/v2.Rand is NOT goroutine-safe, hence the mutex.
type RNG struct {
	mu sync.Mutex
	r  *rand.Rand
}

// New returns an RNG seeded deterministically with seed.
// Same seed → identical sequence (for single-goroutine use).
func New(seed uint64) *RNG {
	return &RNG{r: rand.New(rand.NewPCG(seed, pcgStream))}
}

// Default returns an RNG seeded non-deterministically from crypto/rand
// (falling back to time.Now if crypto/rand fails). Use it when
// reproducibility is not needed.
func Default() *RNG {
	var b [8]byte
	if _, err := crand.Read(b[:]); err == nil {
		return New(binary.LittleEndian.Uint64(b[:]))
	}
	return New(uint64(time.Now().UnixNano()))
}

// Derived returns an RNG whose seed combines seed with the FNV-1a hash
// of domain, so each domain gets an independent stream that is stable
// across runs regardless of call order in other domains:
//
//	Derived(42, "person")   // stable, ≠ internet stream
//	Derived(42, "internet") // stable, ≠ person stream
func Derived(seed uint64, domain string) *RNG {
	h := fnv.New64a()
	_, _ = h.Write([]byte(domain))
	return New(seed ^ h.Sum64())
}

// IntN returns a uniform value in [0, n). It returns 0 for n <= 0
// instead of panicking (math/rand/v2 panics on non-positive n).
func (x *RNG) IntN(n int) int {
	if n <= 0 {
		return 0
	}
	x.mu.Lock()
	defer x.mu.Unlock()
	return x.r.IntN(n)
}

// Int64N returns a uniform value in [0, n). It returns 0 for n <= 0
// instead of panicking.
func (x *RNG) Int64N(n int64) int64 {
	if n <= 0 {
		return 0
	}
	x.mu.Lock()
	defer x.mu.Unlock()
	return x.r.Int64N(n)
}

// Float64 returns a uniform value in [0.0, 1.0).
func (x *RNG) Float64() float64 {
	x.mu.Lock()
	defer x.mu.Unlock()
	return x.r.Float64()
}

// Bool returns a random boolean value.
func (x *RNG) Bool() bool {
	x.mu.Lock()
	defer x.mu.Unlock()
	return x.r.IntN(2) == 0
}

// Shuffle permutes n elements via swap. It is a no-op for n <= 0 or a
// nil swap instead of panicking.
func (x *RNG) Shuffle(n int, swap func(i, j int)) {
	if n <= 0 || swap == nil {
		return
	}
	x.mu.Lock()
	defer x.mu.Unlock()
	x.r.Shuffle(n, swap)
}
