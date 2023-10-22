package random

import (
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Intn returns a random integer in the range [0, n).
func Intn(n int) int {
	return r.Intn(n)
}

// IntnRange returns a random integer in the range [min, max).
func IntnRange(min, max int) int {
	return r.Intn(max-min) + min
}

// Float64 returns a random float64 in the range [0.0, 1.0).
func Float64() float64 {
	return r.Float64()
}
