package main

import (
	"fmt"
)

func test(n1 int) {
	n1++
	fmt.Println("test(), n1 = ", n1)
}

func main() {
	n1 := 10
	test(n1)

	fmt.Println("n1 = ", n1)
}
