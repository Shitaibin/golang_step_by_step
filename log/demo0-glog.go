package main

import (
	"fmt"
	"log"
)

func testLog() string {
	fmt.Println("----------------------------ORZ------------------------")
	return "test log"
}

func main() {
	log.Println(testLog())
}

/*
----------------------------ORZ------------------------
2019/07/11 17:42:12 test log
*/
