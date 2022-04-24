package main

import "fmt"

var age = test()

func test() int {
	fmt.Println("test()函数执行")
	return 90
}

func init() {
	fmt.Println("init函数执行")

}

func main() {

	fmt.Println("main函数执行")
}
