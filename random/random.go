package random

import (
	"math/rand"
	"sync"
	"time"
)

var r struct {
	rand *rand.Rand
	mu   sync.Mutex
}

func init() {
	r = struct {
		rand *rand.Rand
		mu   sync.Mutex
	}{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Seed sets the seed to the provided value.
func Seed(seed int64) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.rand = rand.New(rand.NewSource(seed))
}

// Reseed generates a new seed.
func Reseed() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.rand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Intn returns a random integer in the range [0, n).
func Intn(n int) int {
	if n <= 0 {
		return 0
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.rand.Intn(n)
}

// Float64 returns a random float64 in the range [0.0, 1.0).
func Float64() float64 {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.rand.Float64()
}

// Shuffle shuffles the elements of slice.
func Shuffle[T any](in []T) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i := len(in) - 1; i > 0; i-- {
		j := r.rand.Intn(i + 1)
		in[i], in[j] = in[j], in[i]
	}
}
