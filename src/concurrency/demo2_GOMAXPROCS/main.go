package main

import (
	"fmt"
	"runtime"
)

func main() {

	n := runtime.GOMAXPROCS(1)
	fmt.Printf("%d\n", n)

	//for {
	//	go fmt.Printf("0")
	//	fmt.Printf("1")
	//}

}
