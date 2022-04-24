package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("ccccccccccc")

	//return
	runtime.Goexit()

	fmt.Println("ddddddddddd")
}

func main() {

	go func() {
		fmt.Println("aaaaaaaaaaaaa")

		test()

		fmt.Println("bbbbbbbbbbbbbbb")
	}()

	for {
	}

}
