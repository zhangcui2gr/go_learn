package main

import "fmt"

func test01(arr [3]int) [3]int {
	arr[0] = 10
	arr[1] = 11
	arr[2] = 12
	return arr
}

func test02(arr *[3]int) [3]int {
	(*arr)[0] = 2
	(*arr)[1] = 2
	(*arr)[2] = 3
	return *arr
}

// func test03(arr *[26]byte) [26]byte {
// 	var T byte = 'A'
// 	for index, _ := range *arr {
// 		arr[index] = T + byte(index)
// 	}
// 	return *arr
// }

func test04(arr [5]int) (int, int) {
	max_val, index := arr[0], 0
	for i := 0; i < len(arr); i++ {
		if max_val < arr[i] {
			max_val = arr[i]
			index = i
		}
	}
	return max_val, index
}

func average_arr(arr [5]int) float64 {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return float64(sum) / float64(len(arr))
}

func main() {
	arr := [3]int{1, 2, 3}
	arr2 := test01(arr)
	arr3 := test02(&arr)
	// arr5 := [26]byte{}
	// arr4 := test03(&arr5)
	arr6 := [5]int{1, -1, 9, 90, 10}
	max_val, index := test04(arr6)
	average_num := average_arr(arr6)

	fmt.Println("arr = ", arr)
	fmt.Println("arr2 = ", arr2)
	fmt.Println("arr3 = ", arr3)
	// for i := 0; i < len(arr4); i++ {
	// 	fmt.Printf("%c ", arr4[i])
	// }
	fmt.Println("\n", max_val, index)
	fmt.Println(average_num)
	fmt.Println(5 / 2)

}
