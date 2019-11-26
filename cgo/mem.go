package main

// void SayHello(const char* s);
// char* NewBuf();
// void PrintBuf(char* s);
// void FreeBuf(char* s);
import "C"
import "sync"

// 最大并发数
var maxCnt int
var ch chan int

func main() {
	C.SayHello(C.CString("Dabin\n")) // Just for test

	num := 100000

	maxCnt = 1000
	ch = make(chan int, maxCnt)

	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		id := i
		go mallocPuts(id, &wg)
	}

	wg.Wait()
}

// 使用channel限制实际并发数
func mallocPuts(id int, wg *sync.WaitGroup) {
	ch <- 1
	defer func() {
		<-ch
	}()

	defer wg.Done()

	buf := C.NewBuf()
	defer C.FreeBuf(buf)

	C.PrintBuf(buf)
}
