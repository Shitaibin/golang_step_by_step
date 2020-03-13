package cpucore

import (
	"runtime"
	"sync"
	"testing"
)

const NumOfPlus = 10000
const DataSize = 1000000

// var data [DataSize]int

// 计算密集型任务
// 限制Go routine的数量，完成并发任务
func DoJob(n int) {
	var wg sync.WaitGroup
	gs := make(chan int, n)

	for id := 0; id < DataSize; id++ {
		wg.Add(1)
		gs <- 1
		go Plus(id, &wg, gs)
	}

	// 等待所有g退出
	wg.Wait()
}

func Plus(id int, wg *sync.WaitGroup, gs chan int) {
	num := id
	for i := 0; i < NumOfPlus; i++ {
		num++
	}
	<-gs
	wg.Done()
	// fmt.Printf("%v exit\n", id)
}

func TestDoJob(t *testing.T) {
	DoJob(8)
}

func BenchmarkDoJob8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoJob(8)
	}
}

func BenchmarkDoJob16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoJob(16)
	}
}

/* func BenchmarkDoJob24(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoJob(24)
	}
}

func BenchmarkDoJob32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoJob(32)
	}
}

func BenchmarkDoJob64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoJob(64)
	}
}

func BenchmarkDoJob128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoJob(128)
	}
}

func BenchmarkDoJob512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoJob(512)
	}
}

func BenchmarkDoJob1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoJob(1024)
	}
}

func BenchmarkDoJob2048(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoJob(2048)
	}
}
*/

//--------------------

func BenchmarkDoJobWithP2(b *testing.B) {
	runtime.GOMAXPROCS(2)
	for i := 0; i < b.N; i++ {
		DoJob(8)
	}
}

func BenchmarkDoJobWithP3(b *testing.B) {
	runtime.GOMAXPROCS(3)
	for i := 0; i < b.N; i++ {
		DoJob(8)
	}
}

func BenchmarkDoJobWithP4(b *testing.B) {
	runtime.GOMAXPROCS(4)
	for i := 0; i < b.N; i++ {
		DoJob(8)
	}
}

func BenchmarkDoJobWithP5(b *testing.B) {
	runtime.GOMAXPROCS(5)
	for i := 0; i < b.N; i++ {
		DoJob(8)
	}
}

func BenchmarkDoJobWithP6(b *testing.B) {
	runtime.GOMAXPROCS(6)
	for i := 0; i < b.N; i++ {
		DoJob(8)
	}
}

func BenchmarkDoJobWithP7(b *testing.B) {
	runtime.GOMAXPROCS(7)
	for i := 0; i < b.N; i++ {
		DoJob(8)
	}
}

func BenchmarkDoJobWithP8(b *testing.B) {
	runtime.GOMAXPROCS(8)
	for i := 0; i < b.N; i++ {
		DoJob(8)
	}
}

func BenchmarkDoJobWithP9(b *testing.B) {
	runtime.GOMAXPROCS(9)
	for i := 0; i < b.N; i++ {
		DoJob(8)
	}
}

func BenchmarkDoJobPWith12(b *testing.B) {
	runtime.GOMAXPROCS(12)
	for i := 0; i < b.N; i++ {
		DoJob(8)
	}
}

func BenchmarkDoJobWithP16(b *testing.B) {
	runtime.GOMAXPROCS(16)
	for i := 0; i < b.N; i++ {
		DoJob(8)
	}
}

func BenchmarkDoJobWithP20(b *testing.B) {
	runtime.GOMAXPROCS(20)
	for i := 0; i < b.N; i++ {
		DoJob(8)
	}
}
