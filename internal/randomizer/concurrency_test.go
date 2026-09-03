package randomizer_test

import (
	"sync"
	"testing"

	"github.com/47monad/fakery/internal/randomizer"
)

// TestConcurrentInRange is a regression test for issue #7:
// the shared *rand.Rand was not goroutine-safe. The unseeded path now
// uses the goroutine-safe top-level math/rand/v2 functions, so hammering
// it from many goroutines must pass under -race.
func TestConcurrentInRange(t *testing.T) {
	const goroutines = 50
	const perGoroutine = 10000

	var wg sync.WaitGroup
	for range goroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range perGoroutine {
				v := randomizer.InRange(0, 100)
				if v < 0 || v >= 100 {
					t.Errorf("InRange(0, 100) = %d, want in [0, 100)", v)
					return
				}
			}
		}()
	}
	wg.Wait()
}

// TestConcurrentMixed exercises the other unseeded entry points
// concurrently as well.
func TestConcurrentMixed(t *testing.T) {
	const goroutines = 20
	const perGoroutine = 5000

	words := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	var wg sync.WaitGroup
	for range goroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range perGoroutine {
				_ = randomizer.InRange(0, 100)
				_ = randomizer.Bool()
				_ = randomizer.Digit()
				_ = randomizer.Element(words)
				_ = randomizer.Elements(3, words)
			}
		}()
	}
	wg.Wait()
}
