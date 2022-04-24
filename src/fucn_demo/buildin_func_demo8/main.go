package main

import "fmt"

func main() {

	num1 := 100
	fmt.Printf("num1的地址为 = %v %p \n", &num1, &num1)

	num2 := new(bool)
	fmt.Printf("num2的值为 = %v, 对应的值为 = %v", num2, *num2)

}
