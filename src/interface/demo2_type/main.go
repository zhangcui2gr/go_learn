package main

import "fmt"

type Student struct {
	name string
	id   int
}

func main() {

	t1 := make([]interface{}, 3)
	t1[0] = Student{"sucre", 222}
	t1[1] = "hello"
	t1[2] = 15
	t1 = append(t1, 15.02)
	//if进行类型断言
	for index, val := range t1 {
		if _, ok := val.(int); ok {
			fmt.Printf("第%d个元素类型为int型\n", index)
		} else if _, ok := val.(string); ok {
			fmt.Printf("第%d个元素类型为string型\n", index)
		} else if s, ok := val.(Student); ok {
			fmt.Printf("第%d个元素类型为Student型,姓名为: %s, id为: %v\n", index, s.name, s.id)
		}

	}

	//switch进行类型断言
	for index, val := range t1 {
		switch value := val.(type) {
		case int:
			fmt.Printf("第%d个元素类型为int型\n", index)
		case string:
			fmt.Printf("第%d个元素类型为string型\n", index)
		case Student:
			fmt.Printf("第%d个元素类型为Student型,姓名为: %s, id为: %v\n", index, value.name, value.id)
		default:
			fmt.Printf("第%d个元素类型为其他类型", index)
		}

	}

}
