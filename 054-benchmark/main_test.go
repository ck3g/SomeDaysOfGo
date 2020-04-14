package main

import "testing"

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 2, 3)
	}
}

// To run all benchmarks, use `go test -bench .` or `go test -bench=.`.
