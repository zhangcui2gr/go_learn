package main

import "fmt"

func main() {

	// 方式一：
	arr := [...]int{1, 555, 33, 56, 91}
	slice := arr[:]
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	slice2 := arr[1:]

	fmt.Println("slice2元素 ", slice2)
	fmt.Println("slice元素 ", slice)
	fmt.Println("slice的元素个数 ", len(slice))
	fmt.Println("slice的容量 ", cap(slice))
	fmt.Println("arr的容量是 ", cap(arr))

	// 方式二：
	sli2 := make([]int, 2, 5)
	sli2[0] = 1
	sli2[1] = 3
	for i, v := range sli2 {
		fmt.Printf("i = %v, v = %v\n", i, v)
	}
	fmt.Println("slic2 = ", sli2)

	// 方式三：
	slic3 := []int{0, 1, 2}
	fmt.Println(slic3)

}
