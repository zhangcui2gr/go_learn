package main

import "fmt"

func test01(arr []int) {
	arr[0] = 10
}

func test02() {
	str1 := "hello world!"
	slice1 := str1[5:]
	fmt.Println(slice1)
}

func test03() {
	//string是不可变的，可以将其改成byte数组或[]rune，再转成string
	//[]rune数组可以处理中文和其他特殊字符
	str1 := "hello uestc!"
	fmt.Println(str1)
	arr := []byte(str1)
	arr[0] = 'e'
	str1 = string(arr)
	fmt.Println(str1)

	arr1 := []rune(str1)
	arr1[0] = '中'
	str1 = string(arr1)
	fmt.Println(str1)
}

func main() {
	arr := make([]int, 5)
	fmt.Println(arr)

	test01(arr)
	fmt.Println(arr)

	test02()
	test03()
}
