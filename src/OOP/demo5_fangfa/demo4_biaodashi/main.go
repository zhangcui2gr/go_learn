package main

import "fmt"

type Student struct {
	Class, Age int
	Name       string
}

func (S *Student) Myprint() {
	fmt.Println("调用了Myprint函数!")
}

func (S Student) Myprint2() {
	fmt.Println("调用了Myprint2函数!")
}

func main() {

	s1 := Student{3, 15, "xiaoguan"}

	//方法值
	f1 := s1.Myprint
	f1()

	//方法表达式
	f2 := (*Student).Myprint
	f2(&s1)

	f3 := Student.Myprint2
	f3(s1)
}
