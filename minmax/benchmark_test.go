package minmax

import "testing"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// 从testing/benchmark.go看到了这种奇怪的写法
// 逻辑上来讲，性能应没有差异。
// 测试结果也是如此。
func goMin(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func goMax(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func BenchmarkMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		min(1, 2)
	}
}

func BenchmarkGoMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goMin(1, 2)
	}
}

func BenchmarkMinParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			min(1, 2)
		}
	})
}

func BenchmarkGoMinParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			goMin(1, 2)
		}
	})
}
