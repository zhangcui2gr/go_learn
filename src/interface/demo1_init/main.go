package main

import "fmt"

type Humaner interface {
	test()
}

type Human struct {
	Name string
	Id   int
}

func (h *Human) test() {
	fmt.Printf("Human的Name: %s, Id: %d\n", h.Name, h.Id)
}

type Student struct {
	Name  string
	class int
}

func (s *Student) test() {
	fmt.Printf("Student的Name: %s, class: %d\n", s.Name, s.class)
}

type Mystr string

func (tmp *Mystr) test() {
	fmt.Printf("Mystr: %v", *tmp)
}

//普通函数实现接口的封装
func test_call(tmp Humaner) {
	tmp.test()
}

func main() {

	h1 := Human{"xiaoguan", 123}
	s1 := Student{"xiaozhang", 21}
	var str1 Mystr = "xiaocheng"

	test_call(&h1)
	test_call(&s1)
	test_call(&str1)

}
