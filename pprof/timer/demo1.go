// package timer_demo

// func test() {
// 	n := 10
// 	var wg sync.WaitGroup

// 	wait := func() {
// 		defer wg.Done()
// 		// defer fmt.Println("exit")

// 		ch := time.Tick(time.Second)
// 		<-ch
// 	}

// 	wg.Add(n)
// 	for i := 0; i < n; i++ {
// 		go wait()
// 	}

// 	wg.Wait()
// 	// fmt.Println("All exit")
// }

// func main() {
// 	test()
// }
