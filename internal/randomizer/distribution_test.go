package randomizer_test

import (
	"math/rand/v2"
	"testing"

	"github.com/47monad/fakery/internal/randomizer"
)

func TestDigitDistribution(t *testing.T) {
	draws := 10000
	seen := make(map[int]bool)
	for range draws {
		d := randomizer.Digit()
		if d < 0 || d > 9 {
			t.Fatalf("Digit() = %d, want a digit in [0, 9]", d)
		}
		seen[d] = true
	}
	for d := range 10 {
		if !seen[d] {
			t.Fatalf("digit %d never appeared in %d draws", d, draws)
		}
	}
}

func TestBoolDistribution(t *testing.T) {
	draws := 10000
	trues, falses := 0, 0
	for range draws {
		if randomizer.Bool() {
			trues++
		} else {
			falses++
		}
	}
	if trues == 0 || falses == 0 {
		t.Fatalf("expected both true and false in %d draws, got trues=%d falses=%d", draws, trues, falses)
	}
	// A fair coin should stay well within 40-60% over 10k draws.
	ratio := float64(trues) / float64(draws)
	if ratio < 0.4 || ratio > 0.6 {
		t.Fatalf("unbalanced distribution: trues=%d falses=%d", trues, falses)
	}
}

func TestBoolWithDistribution(t *testing.T) {
	r := rand.New(rand.NewPCG(42, 17))
	trues, falses := 0, 0
	for range 10000 {
		if randomizer.BoolWith(r) {
			trues++
		} else {
			falses++
		}
	}
	if trues == 0 || falses == 0 {
		t.Fatalf("expected both true and false, got trues=%d falses=%d", trues, falses)
	}
}

func TestNumberBounds(t *testing.T) {
	tests := []struct {
		length uint32
		low    int
		high   int
	}{
		{2, 10, 99},
		{5, 10000, 99999},
		{9, 100000000, 999999999},
	}
	for _, tt := range tests {
		for range 1000 {
			n := randomizer.Number(tt.length)
			if n < tt.low || n > tt.high {
				t.Fatalf("Number(%d) = %d, want in [%d, %d]", tt.length, n, tt.low, tt.high)
			}
		}
	}
}

func TestElementsClampsCount(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	got := randomizer.Elements(10, slice)
	if len(got) != len(slice) {
		t.Fatalf("Elements(10, ...) returned %d elements, want clamped to %d", len(got), len(slice))
	}

	seen := make(map[int]bool)
	for _, v := range got {
		seen[v] = true
	}
	for _, v := range slice {
		if !seen[v] {
			t.Fatalf("Elements(10, ...) missed element %d", v)
		}
	}
}

func TestElementsEmptySlice(t *testing.T) {
	if got := randomizer.Elements(3, []int{}); len(got) != 0 {
		t.Fatalf("Elements on empty slice returned %v, want empty", got)
	}
}

func TestElementsDistinct(t *testing.T) {
	slice := []string{"a", "b", "c", "d", "e", "f"}
	got := randomizer.Elements(4, slice)
	if len(got) != 4 {
		t.Fatalf("Elements(4, ...) returned %d elements, want 4", len(got))
	}
	seen := make(map[string]bool)
	for _, v := range got {
		if seen[v] {
			t.Fatalf("Elements returned duplicate element %q", v)
		}
		seen[v] = true
	}
}
