package main

var c = make(chan int)
var a int

func f() {
	a = 1
	<-c
}
func main() {
	go f()
	c <- 0
	print(a)
}
