package main

import "fmt"

func test01() {
	monsters := make([]map[string]string, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 5)
		monsters[0]["name"] = "科加斯"
		monsters[0]["age"] = "200"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 5)
		monsters[1]["name"] = "阿木木"
		monsters[1]["age"] = "100"
	}

	for key, val := range monsters {
		fmt.Println(key)
		for k, v := range val {
			fmt.Printf("\t k=%v v = %v\n", k, v)
		}
	}

}

func main() {

	test01()

}
