package main

import "fmt"

type Person struct {
	Age, Id int
	name    string
}

func (p *Person) print_val() {
	fmt.Println("Person的print_val函数调用!")
}

type Student struct {
	Person
	class, score int
}

func (s *Student) print_val() {
	fmt.Println("Student的print_val函数调用!")
}

func main() {

	//就近原则，调用自己作用域的方法或字段
	p1 := Person{26, 252, "xiaozhang"}
	p1.print_val()

	s1 := Student{Person{25, 123, "xiaoguan"}, 2, 95}
	s1.print_val()

	//显示的调用
	s1.Person.print_val()

}
