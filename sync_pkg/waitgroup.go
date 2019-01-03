package main

import (
	"fmt"
	"sync"
)

func main() {
	leader()
}

func leader() {
	var wg sync.WaitGroup
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go follower(&wg, i)
	}
	wg.Wait()
	
	fmt.Println("open the box together")
}

func follower(wg *sync.WaitGroup, id int) {
	fmt.Printf("follwer %d find key\n", id)
	wg.Done()
}
