package randomizer_test

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand/v2"
	"testing"
	"time"

	"github.com/47monad/fakery/internal/randomizer"
)

func TestCollision(t *testing.T) {
	count := 1000
	max := 10000
	seen := make(map[int]bool)
	for i := range count {
		num := randomizer.InRange(1, max)
		if seen[num] {
			fmt.Printf("collision occured at %d with default rng \n", i+1)
			break
		} else {
			seen[num] = true
		}
	}

	seen2 := make(map[int]bool)
	rand2 := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 17))
	for i := range count {
		num := randomizer.InRangeWith(rand2, 1, max)
		if seen2[num] {
			fmt.Printf("collision occured at %d with provided time rng \n", i+1)
			break
		} else {
			seen2[num] = true
		}
	}

	seen3 := make(map[int]bool)
	var seed uint64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &seed); err != nil {
		t.Fatalf("failed to read crypto seed: %v", err)
	}
	rand3 := rand.New(rand.NewPCG(seed, 17))
	for i := range count {
		num := randomizer.InRangeWith(rand3, 1, max)
		if seen3[num] {
			fmt.Printf("collision occured at %d with provided crypto rng \n", i+1)
			break
		} else {
			seen3[num] = true
		}
	}
}
