package main

import "fmt"

type Person struct {
	Name string
	Age  int
	sex  string
}

type Student struct {
	Person
	Id   int
	addr string
}

func main() {
	//顺序初始化
	stu01 := Student{Person{"xiaoming", 15, "男"}, 123, "SH"}
	fmt.Println(stu01)

	//部分初始化
	stu02 := Student{Person: Person{"xiaoguan", 25, "男"}, Id: 122}
	fmt.Println(stu02)

	//部分初始化
	stu03 := Student{Person: Person{Name: "xiaozhang", Age: 26}, addr: "CD"}
	fmt.Println(stu03)

	//Person结构中的成员也归属于Student中了
	var stu04 Student
	stu04.Age = 23
	fmt.Println(stu04)
	//也可以直接对Person进行初始化
	var stu05 Student
	stu05.Person = Person{Name: "xiaohuang"}
	fmt.Println(stu05)
	fmt.Printf("%+v\n", stu05)
	fmt.Printf("%v\n", stu05)

}
