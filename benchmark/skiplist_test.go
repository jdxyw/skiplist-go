package benchmark

import (
	"math/rand"
	"skiplist-go"
	"testing"
)

const fixedByetesLen = 6

func benchmark(level, n int) {
	s := skiplist.NewSkiplist(level, nil)
	rand.Seed(6)
	keys := make([][]byte, 0)
	for i := 0; i < n; i++ {
		key := make([]byte, fixedByetesLen)
		rand.Read(key)
		s.Set(key, nil)

		if rand.Float32() < 0.05 {
			keys = append(keys, key)
		}
	}

	for _, key := range keys {
		s.Get(key)
	}

	for _, key := range keys {
		s.Delete(key)
	}
}

func BenchmarkLevel6(b *testing.B) {
	benchmark(6, b.N)
}

func BenchmarkLevel8(b *testing.B) {
	benchmark(8, b.N)
}

func BenchmarkLevel10(b *testing.B) {
	benchmark(10, b.N)
}

func BenchmarkLevel12(b *testing.B) {
	benchmark(12, b.N)
}
