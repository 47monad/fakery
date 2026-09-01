package fkrand_test

import (
	"testing"

	"github.com/47monad/fakery/fk/fkrand"
)

func TestDigitDistribution(t *testing.T) {
	draws := 10000
	seen := make(map[int]bool)
	for range draws {
		d := fkrand.Digit()
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
	trues, falses := 0, 0
	for range 10000 {
		if fkrand.Bool() {
			trues++
		} else {
			falses++
		}
	}
	if trues == 0 || falses == 0 {
		t.Fatalf("expected both true and false, got trues=%d falses=%d", trues, falses)
	}
}

func TestIntLengths(t *testing.T) {
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
			n := fkrand.Int(tt.length)
			if n < tt.low || n > tt.high {
				t.Fatalf("Int(%d) = %d, want in [%d, %d]", tt.length, n, tt.low, tt.high)
			}
		}
	}
}

func TestIntTooLargePanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("Int(10) should panic")
		}
	}()
	fkrand.Int(10)
}

func TestElementsClampsCount(t *testing.T) {
	slice := []int{1, 2, 3}
	got := fkrand.Elements(100, slice)
	if len(got) != len(slice) {
		t.Fatalf("Elements(100, ...) returned %d elements, want clamped to %d", len(got), len(slice))
	}
}
