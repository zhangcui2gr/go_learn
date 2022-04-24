package main

import (
	"fmt"
	"sort"
)

func test01() {

	// slice2 := make([]int, 1)

	m1 := make(map[int]int)
	m1[0] = 2
	m1[3] = 1
	m1[2] = 4
	m1[100] = 20
	fmt.Println(m1)
	for k, v := range m1 {
		fmt.Printf("k = %v, v = %v\n", k, v)
	}

	var slice1 []int
	for k, v := range m1 {
		slice1 = append(slice1, k)
		fmt.Println(v)
	}
	sort.Ints(slice1)
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("k = %v, v = %v \n", slice1[i], m1[slice1[i]])
	}

}

func main() {

	test01()

}
