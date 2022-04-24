package main

import "fmt"

var (
	Fun = func(n1, n2 int) int {
		return n1 * n2
	}
)

func main() {

	res := func(n1, n2 int) int {
		fmt.Println("调用了匿名函数!")
		return n1 + n2
	}

	// var num1 float32 = 2

	fmt.Println(res(20, 30))
	fmt.Printf("%T\n", res)

	fmt.Println(Fun(20, 30))
	fmt.Printf("%T\n", res)
}
