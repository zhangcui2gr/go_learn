package main

import (
	"fmt"
	"unsafe"
)

type cat struct {
	Color  string
	Age    int64
	weight float64
	slic1  []int
	p      *int
}

func main() {

	var cat1 cat
	cat1.weight = 10
	cat1.Age = 2
	cat1.Color = "black"
	cat1.slic1 = make([]int, 2)
	cat1.slic1[0] = 1
	cat1.p = new(int)

	fmt.Println(cat1)
	fmt.Printf("%v \n", unsafe.Sizeof(cat1))

	str1 := "name"

	fmt.Printf("%v \n", unsafe.Sizeof(str1))

	p1 := new(int)
	slice2 := make([]int, 2)
	// cat2 := cat{"red", 1, 4.2, slice2, p1}
	//指定成员初始化
	cat2 := cat{p: p1, slic1: slice2}
	fmt.Println(cat2)

	cat3 := &cat{Age: 1}
	fmt.Printf("type = %T, val = %v", cat3, *cat3)

}
