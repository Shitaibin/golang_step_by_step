package test

import "testing"

// go test -bench=. -benchmem ./...

// 样例函数
// 两种交互写法性能并无显著差别
func swap(i, j int) {
	i, j = j, i
}

func swapT(i, j int) {
	t := i
	i = j
	j = t
}

// 测试接口非并发时性能

func BenchmarkSwap(b *testing.B) {
	// 不要省略这个for，不然就不准了
	for i := 0; i < b.N; i++ {
		swap(1, 2)
	}
}

func BenchmarkSwapT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		swapT(1, 2)
	}
}

// 测试接口的并发性能

func BenchmarkSwapParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			swap(1, 2)
		}
	})
}

func BenchmarkSwapTParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			swapT(1, 2)
		}
	})
}

// 使用-v时，可以看到运行benchmark，会先运行test，可以使用-run=^$过滤掉所有test
func TestSwap(t *testing.T) {
	swap(1, 2)
}
