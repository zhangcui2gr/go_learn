package main

import "fmt"

func test() {
	arr := [2][3]int8{}
	fmt.Printf("%p \n", &arr[0])
	fmt.Printf("%p", &arr[1])

}

func main() {

	test()

}
