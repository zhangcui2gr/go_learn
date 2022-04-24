package main

import (
	"fmt"
	"math/rand"
	"time"
)

func test() {
	var arr [6]float64
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Float64()
		fmt.Printf("第 %d 个数是: %f\n", i+1, arr[i])
	}
	// fmt.Printf("%T", arr)
	for index, val := range arr {
		fmt.Printf("第 %d 个是: %.2f \n", index+1, val)
	}

	arr2 := [3]int{1, 3, 2}
	fmt.Println(arr2)
	var num1 = [...]int{0: 4, 2: 5}
	fmt.Println(num1)

	var arr3 []int
	// arr3[0] = 1
	// arr3[1] = 2
	fmt.Printf("%T \n", arr3)

}

func main() {

	test()
	fmt.Println("main函数执行到这了!")
}
