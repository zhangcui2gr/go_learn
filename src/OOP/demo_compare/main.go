package main

import "fmt"

type student struct {
	Id, Age    int
	name, addr string
}

func deliver(stu student) {
	stu.Id = 11
}

func deliver_arr(stu *student) {
	stu.Age = 16
}

func main() {

	s1 := student{12, 15, "mike", "bj"}
	s2 := student{12, 15, "mike", "bj"}
	s3 := student{12, 15, "mike", "sh"}

	f1 := s1 == s2
	f2 := s2 == s3

	fmt.Println("f1 = ", f1, "f2 = ", f2)

	fmt.Println(s3)
	deliver(s3)
	fmt.Println(s3)

	deliver_arr(&s3)
	fmt.Println(s3)

}
