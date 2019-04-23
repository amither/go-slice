package main

import (
	"testing"
)

func BenchmarkCo(b *testing.B) {
	s := []byte{'1', '2', '3', '4'}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PermutationConcurrency(s)
	}
}

func BenchmarkSi(b *testing.B) {
	s := []byte{'1', '2', '3', '4'}
	p := make([]byte, len(s), len(s))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		permutaionImpl(p, s, 0)
	}
}
