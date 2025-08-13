package randomizer

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand/v2"
	"sync"
	"time"
)

var (
	defaultRNG *rand.Rand
	rngMu      sync.Mutex
	once       sync.Once
)

type RNGSource interface {
	NewRand() *rand.Rand
}

type SeedSource struct {
	seed uint64
}

func (s *SeedSource) NewRand() *rand.Rand {
	return rand.New(rand.NewPCG(s.seed, 17))
}

type TimeSource struct{}

func (t *TimeSource) NewRand() *rand.Rand {
	return rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 17))
}

type CryptoSource struct{}

func (t *CryptoSource) NewRand() *rand.Rand {
	var seed int64
	binary.Read(crand.Reader, binary.LittleEndian, &seed)
	return rand.New(rand.NewPCG(uint64(seed), 17))
}

func getDefaultRNG() *rand.Rand {
	once.Do(func() {
		defaultRNG = new(TimeSource).NewRand()
	})
	return defaultRNG
}
