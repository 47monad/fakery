package rng_test

import (
	"sync"
	"testing"

	"github.com/47monad/fakery/internal/rng"
)

func drainSequence(r *rng.RNG, n int) []uint64 {
	// Mix every method so determinism covers the whole surface,
	// including Shuffle ordering.
	out := make([]uint64, 0, n*4)
	for range n {
		out = append(out,
			uint64(r.IntN(1_000_000)),
			uint64(r.Int64N(1_000_000)),
		)
		out = append(out, uint64(r.Float64()*1e9))
		if r.Bool() {
			out = append(out, 1)
		} else {
			out = append(out, 0)
		}
	}
	perm := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r.Shuffle(len(perm), func(i, j int) { perm[i], perm[j] = perm[j], perm[i] })
	for _, v := range perm {
		out = append(out, uint64(v))
	}
	return out
}

func equal(a, b []uint64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Same seed → identical 10k-value sequence. Re-running the test binary
// (or the process) must produce the same result, which is what makes
// this a cross-run determinism check.
func TestDeterminism(t *testing.T) {
	const n = 10000
	a := drainSequence(rng.New(42), n)
	b := drainSequence(rng.New(42), n)
	if !equal(a, b) {
		t.Fatal("same seed produced different sequences")
	}
	c := drainSequence(rng.New(43), n)
	if equal(a, c) {
		t.Fatal("different seeds produced identical sequences")
	}
}

// Concurrent use of one RNG must be safe under -race. Ordering is
// intentionally nondeterministic here; we only assert safety and bounds.
func TestConcurrentUse(t *testing.T) {
	r := rng.New(1)
	const goroutines = 50
	const perGoroutine = 2000

	var wg sync.WaitGroup
	for range goroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range perGoroutine {
				if v := r.IntN(100); v < 0 || v >= 100 {
					t.Errorf("IntN(100) = %d, want in [0, 100)", v)
					return
				}
				if v := r.Int64N(100); v < 0 || v >= 100 {
					t.Errorf("Int64N(100) = %d, want in [0, 100)", v)
					return
				}
				_ = r.Float64()
				_ = r.Bool()
				r.Shuffle(5, func(i, j int) {})
			}
		}()
	}
	wg.Wait()
}

// Derived streams for different domains must differ, but each must be
// stable across runs.
func TestDerived(t *testing.T) {
	const n = 5000
	personA := drainSequence(rng.Derived(42, "person"), n)
	personB := drainSequence(rng.Derived(42, "person"), n)
	internet := drainSequence(rng.Derived(42, "internet"), n)

	if !equal(personA, personB) {
		t.Fatal("Derived(42, \"person\") is not stable across runs")
	}
	if equal(personA, internet) {
		t.Fatal("Derived(42, \"person\") == Derived(42, \"internet\"), want different streams")
	}

	emptyA := drainSequence(rng.Derived(42, ""), n)
	emptyB := drainSequence(rng.Derived(42, ""), n)
	if !equal(emptyA, emptyB) {
		t.Fatal("Derived with empty domain is not stable")
	}
}

func TestNoPanics(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("panicked: %v", r)
		}
	}()
	r := rng.New(0)
	_ = r.IntN(0)
	_ = r.IntN(-5)
	_ = r.Int64N(0)
	_ = r.Int64N(-5)
	_ = r.Float64()
	_ = r.Bool()
	r.Shuffle(0, nil)
	r.Shuffle(-3, nil)
	r.Shuffle(5, nil)
	r.Shuffle(0, func(i, j int) {})
	_ = rng.Derived(0, "")
	_ = rng.Default()
}
