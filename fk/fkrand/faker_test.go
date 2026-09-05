package fkrand_test

import (
	"sync"
	"testing"

	"github.com/47monad/fakery/fk/fkrand"
)

// drain pulls one value from every generator so determinism covers the
// whole surface.
func drain(f *fkrand.Faker) []any {
	words := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	return []any{
		f.Digit(),
		f.Int(5),
		f.IntInRange(-10, 10),
		f.Float64(),
		f.Float64InRange(-2.5, 7.5),
		f.Bool(),
		fkrand.Element(words, f),
		fkrand.Elements(3, words, f),
	}
}

func drainMany(f *fkrand.Faker, n int) [][]any {
	out := make([][]any, 0, n)
	for range n {
		out = append(out, drain(f))
	}
	return out
}

func equalSequences(a, b [][]any) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			sa, sb := a[i][j], b[i][j]
			// []string needs element-wise compare; the rest is comparable.
			la, oka := sa.([]string)
			lb, okb := sb.([]string)
			if oka && okb {
				if len(la) != len(lb) {
					return false
				}
				for k := range la {
					if la[k] != lb[k] {
						return false
					}
				}
				continue
			}
			if sa != sb {
				return false
			}
		}
	}
	return true
}

func TestSeededDeterminism(t *testing.T) {
	a := drainMany(fkrand.New(fkrand.WithSeed(42)), 500)
	b := drainMany(fkrand.New(fkrand.WithSeed(42)), 500)
	if !equalSequences(a, b) {
		t.Fatal("same seed produced different sequences")
	}
	c := drainMany(fkrand.New(fkrand.WithSeed(43)), 500)
	if equalSequences(a, c) {
		t.Fatal("different seeds produced identical sequences")
	}
}

func TestSeededDigitDistribution(t *testing.T) {
	f := fkrand.New(fkrand.WithSeed(7))
	seen := make(map[int]bool)
	for range 10000 {
		d := f.Digit()
		if d < 0 || d > 9 {
			t.Fatalf("Digit() = %d, want in [0, 9]", d)
		}
		seen[d] = true
	}
	for d := range 10 {
		if !seen[d] {
			t.Fatalf("digit %d never appeared (regression for #8)", d)
		}
	}
}

func TestSeededIntInRangeHalfOpen(t *testing.T) {
	f := fkrand.New(fkrand.WithSeed(7))
	for range 10000 {
		v := f.IntInRange(3, 8)
		if v < 3 || v >= 8 {
			t.Fatalf("IntInRange(3, 8) = %d, want in [3, 8)", v)
		}
	}
	seenMax := false
	for range 5000 {
		if v := f.IntInRange(0, 1); v != 0 {
			t.Fatalf("IntInRange(0, 1) = %d, want always 0", v)
		}
		if v := f.IntInRange(5, 6); v != 5 {
			t.Fatalf("IntInRange(5, 6) = %d, want always 5", v)
		}
		if f.IntInRange(0, 100) == 100 {
			seenMax = true
		}
	}
	if seenMax {
		t.Fatal("IntInRange(0, 100) returned 100, want half-open [min, max)")
	}
}

func TestSeededIntLengths(t *testing.T) {
	f := fkrand.New(fkrand.WithSeed(7))
	tests := []struct {
		length uint32
		low    int
		high   int
	}{
		{0, 0, 9},
		{1, 0, 9},
		{2, 10, 99},
		{5, 10000, 99999},
		{9, 100000000, 999999999},
	}
	for _, tt := range tests {
		for range 1000 {
			n := f.Int(tt.length)
			if n < tt.low || n > tt.high {
				t.Fatalf("Int(%d) = %d, want in [%d, %d]", tt.length, n, tt.low, tt.high)
			}
		}
	}
}

func TestSeededIntTooLargePanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("Int(10) should panic")
		}
	}()
	fkrand.New(fkrand.WithSeed(1)).Int(10)
}

func TestSeededFloatRanges(t *testing.T) {
	f := fkrand.New(fkrand.WithSeed(7))
	for range 5000 {
		if v := f.Float64(); v < 0.0 || v >= 1.0 {
			t.Fatalf("Float64() = %v, want in [0.0, 1.0)", v)
		}
		if v := f.Float64InRange(2.0, 5.0); v < 2.0 || v >= 5.0 {
			t.Fatalf("Float64InRange(2, 5) = %v, want in [2.0, 5.0)", v)
		}
		if v := f.Float64InRange(-5.0, -1.0); v < -5.0 || v >= -1.0 {
			t.Fatalf("Float64InRange(-5, -1) = %v, want in [-5.0, -1.0)", v)
		}
	}
	if v := f.Float64InRange(4.0, 4.0); v != 4.0 {
		t.Fatalf("Float64InRange(4, 4) = %v, want 4.0", v)
	}
	if v := f.Float64InRange(5.0, 4.0); v != 5.0 {
		t.Fatalf("Float64InRange(5, 4) = %v, want min 5.0", v)
	}
}

func TestSeededElementHitsEveryIndex(t *testing.T) {
	f := fkrand.New(fkrand.WithSeed(7))
	words := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	seen := make(map[string]bool)
	for range 10000 {
		seen[fkrand.Element(words, f)] = true
	}
	for _, w := range words {
		if !seen[w] {
			t.Fatalf("Element never picked %q in 10000 draws", w)
		}
	}
}

func TestSeededElementEmptySliceReturnsZero(t *testing.T) {
	f := fkrand.New(fkrand.WithSeed(7))
	if v := fkrand.Element([]int{}, f); v != 0 {
		t.Fatalf("Element(empty) = %d, want zero value", v)
	}
	if v := fkrand.Element([]string{}, f); v != "" {
		t.Fatalf("Element(empty) = %q, want zero value", v)
	}
	if v := fkrand.Element[int](nil, f); v != 0 {
		t.Fatalf("Element(nil) = %d, want zero value", v)
	}
	// Tier-3 path must not panic either.
	if v := fkrand.Element([]int{}); v != 0 {
		t.Fatalf("tier-3 Element(empty) = %d, want zero value", v)
	}
}

func TestSeededElementsClampAndDistinct(t *testing.T) {
	f := fkrand.New(fkrand.WithSeed(7))
	slice := []int{1, 2, 3, 4, 5}

	got := fkrand.Elements(100, slice, f)
	if len(got) != len(slice) {
		t.Fatalf("Elements(100, ...) returned %d, want clamped to %d", len(got), len(slice))
	}
	seen := make(map[int]bool)
	for _, v := range got {
		if seen[v] {
			t.Fatalf("Elements returned duplicate %d", v)
		}
		seen[v] = true
	}
	for _, v := range slice {
		if !seen[v] {
			t.Fatalf("clamped Elements missed %d", v)
		}
	}

	if got := fkrand.Elements(0, slice, f); len(got) != 0 {
		t.Fatalf("Elements(0, ...) = %v, want empty", got)
	}
	if got := fkrand.Elements(-3, slice, f); len(got) != 0 {
		t.Fatalf("Elements(-3, ...) = %v, want empty", got)
	}
	if got := fkrand.Elements(3, []int{}, f); len(got) != 0 {
		t.Fatalf("Elements on empty slice = %v, want empty", got)
	}
	if got := fkrand.Elements[int](3, nil, f); len(got) != 0 {
		t.Fatalf("Elements on nil slice = %v, want empty", got)
	}
}

func TestSeededNoPanicsOnExtremeInputs(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("panicked on extreme inputs: %v", r)
		}
	}()
	f := fkrand.New(fkrand.WithSeed(0))
	_ = f.IntInRange(-100, -100) // min == max → min
	_ = f.IntInRange(10, 5)      // reversed → min
	_ = f.IntInRange(-50, -40)   // negative range
	_ = f.Float64InRange(1.0, 1.0)
	_ = f.Float64InRange(2.0, 1.0)
	_ = f.Float64InRange(-3.0, -3.0)
	_ = fkrand.Element([]int{}, f)
	_ = fkrand.Element[int](nil, f)
	_ = fkrand.Elements(5, []int{}, f)
	_ = fkrand.Elements(-1, []int{1}, f)
	_ = fkrand.Elements[int](3, nil, f)
	// Nil-faker tolerance.
	_ = fkrand.Element([]int{1, 2, 3}, nil)
	_ = fkrand.Elements(2, []int{1, 2, 3}, nil)
	var nilFaker *fkrand.Faker
	_ = nilFaker.Digit()
	_ = nilFaker.Int(2)
	_ = nilFaker.IntInRange(0, 10)
	_ = nilFaker.Float64()
	_ = nilFaker.Float64InRange(0, 1)
	_ = nilFaker.Bool()
	// Tier-3 smoke.
	_ = fkrand.Digit()
	_ = fkrand.Int(2)
	_ = fkrand.IntInRange(-5, 5)
	_ = fkrand.Float64()
	_ = fkrand.Float64InRange(-1, 1)
	_ = fkrand.Bool()
	_ = fkrand.Element([]string{"x"})
	_ = fkrand.Elements(2, []string{"x", "y", "z"})
}

func TestSeededEdgeValues(t *testing.T) {
	f := fkrand.New(fkrand.WithSeed(1))
	if v := f.IntInRange(7, 7); v != 7 {
		t.Fatalf("IntInRange(7, 7) = %d, want 7", v)
	}
	if v := f.IntInRange(9, 4); v != 9 {
		t.Fatalf("IntInRange(9, 4) = %d, want min 9", v)
	}
}

func TestSeededConcurrentUse(t *testing.T) {
	f := fkrand.New(fkrand.WithSeed(1))
	words := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	var wg sync.WaitGroup
	for range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 2000 {
				_ = f.Digit()
				_ = f.Int(3)
				_ = f.IntInRange(0, 100)
				_ = f.Float64()
				_ = f.Float64InRange(0, 100)
				_ = f.Bool()
				_ = fkrand.Element(words, f)
				_ = fkrand.Elements(3, words, f)
			}
		}()
	}
	wg.Wait()
}

func TestTier3Bounds(t *testing.T) {
	for range 1000 {
		if d := fkrand.Digit(); d < 0 || d > 9 {
			t.Fatalf("tier-3 Digit() = %d", d)
		}
		if v := fkrand.IntInRange(0, 10); v < 0 || v >= 10 {
			t.Fatalf("tier-3 IntInRange(0, 10) = %d", v)
		}
		if v := fkrand.Float64(); v < 0.0 || v >= 1.0 {
			t.Fatalf("tier-3 Float64() = %v", v)
		}
		if v := fkrand.Float64InRange(1.0, 2.0); v < 1.0 || v >= 2.0 {
			t.Fatalf("tier-3 Float64InRange(1, 2) = %v", v)
		}
	}
}
