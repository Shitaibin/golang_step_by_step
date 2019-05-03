package doublecheck

import (
	"testing"
)

func BenchmarkUnsafety(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Testunsafety()
	}
}

func BenchmarkSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TestSlow()
	}
}

func BenchmarkFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TestFast()
	}
}
