package main

import "fmt"

//普通函数实现
func My_add1(a, b float64) float64 {
	return a + b
}

type num float64

//方法实现
func (num1 num) My_add2(num2 num) num {
	return num1 + num2
}

func main() {

	a := num(5)
	b := num(6.1)

	fmt.Println(My_add1(1, 2))
	fmt.Println(a.My_add2(b))

}
