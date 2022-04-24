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
	Name string
}

type Student01 struct {
	*Person
	Id   int
	addr string
	Name string
}

func main() {

	// stu01 := Student{Person: Person{Name: "xiaoming"}}
	var stu01 Student
	stu01.Name = "xiaozhang1"
	fmt.Printf("%+v\n", stu01)

	var stu02 Student01
	stu02.Person = new(Person)
	stu02.Person.Age = 15
	fmt.Printf("%+v\t %+v\n", stu02, *(stu02.Person))

}
