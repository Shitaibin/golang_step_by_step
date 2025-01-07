package main

import "fmt"

func main() {
	test()
	test2()
}

func test() (msg string) {
	defer func() {
		fmt.Printf("catch: %s\n", msg)
	}()

	if msg := genMsg(); len(msg) != 0 {
		fmt.Printf("msg: %s\n", msg)
		return msg
	}
	return "no valid msg"
}

func genMsg() string {
	return "hello"
}

func test2() {
	fmt.Println("---------")

	var err error
	defer func(){
		if err != nil{
			fmt.Printf("catch error: %s\n", err.Error())
		}
	}()

	err = fmt.Errorf("123")
	return
}