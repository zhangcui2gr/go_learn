package main

import "fmt"

var Age int = 20

func sum(n1, n2 int) int {
	defer fmt.Println("ok n1 = ", n1)
	defer fmt.Println("ok n2 = ", n2)

	n1++
	res := n1 + n2

	fmt.Println("ok res = ", res)
	return res
}

func main() {

	res := sum(10, 20)

	fmt.Println("res = ", res)

}
