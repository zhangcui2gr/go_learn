package main

import "fmt"

type Person struct {
	Age, Id int
}

//给实例化的某个Person对象定义方法
func (p Person) My_print() {
	fmt.Printf("Age = %v, Id = %v\n", p.Age, p.Id)
}

//接受者为指针变量，此时为引用语义
func (p *Person) Setinfo_pointer(a, b int) {
	p.Age = a
	p.Id = b
}

//接收者为普通变量，此时为值语义
func (p Person) Setinfo_val(a, b int) {
	p.Age = a
	p.Id = b
}

func main() {

	p1 := Person{24, 234}
	p1.My_print()
	p1.Setinfo_val(50, 321)
	fmt.Printf("p1 = %+v\n", p1)

	var p2 Person
	p2.Setinfo_pointer(30, 254)

	var p3 Person
	p3.Setinfo_pointer(25, 123)

	fmt.Printf("p2 = %+v\np3 = %+v", p2, p3)

}
