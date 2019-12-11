package test

import "testing"

func fast(i, j int) {
	i, j = j, i
}

func slow(i, j int) {
	t := i
	i = j
	j = t
}

func BenchmarkFast(b *testing.B) {
	// 不要省略这个for，不然就不准了
	for i := 0; i < b.N; i++ {
		fast(1, 2)
	}
}

func BenchmarkSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slow(1, 2)
	}
}
