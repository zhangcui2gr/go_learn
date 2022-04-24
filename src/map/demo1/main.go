package main

import "fmt"

func test01() {
	// var m map[int]string
	m := make(map[int]string, 1)

	m[1] = "宋江"
	m[2] = "宋江"
	m[3] = "宋江"

	fmt.Println(m)
}

func test02() {
	// var m1 map[int]int
	m1 := make(map[int]int, 2)
	m1[0] = 10
	fmt.Println(m1)

	m2 := make(map[int]int, 3)
	m2[1] = 20
	fmt.Println(m2)

	var m3 map[int]int = map[int]int{
		1: 1,
	}
	m3[2] = 2
	fmt.Println(m3)
}

func test03() {
	m1 := make(map[string]map[string]string)
	m1["stu01"] = make(map[string]string)

	m1["stu01"]["name"] = "小管"
	m1["stu01"]["age"] = "25"

	fmt.Println(m1)

	delete(m1["stu01"], "name")
	delete(m1, "stu02")
	fmt.Println(m1)

	// m1 = make(map[string]map[string]string)
	// fmt.Println(m1)

	val, ok := m1["stu02"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("找不到")
	}

	m1["stu01"]["height"] = "178"
	m1["stu01"]["weight"] = "125"

	fmt.Println(m1)

	for _, val := range m1 {
		for _, val1 := range val {
			fmt.Println(val1)
		}
	}

	fmt.Println(len(m1["stu01"]))

}

func main() {

	test01()

	test02()

	test03()

}
