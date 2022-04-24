package main

import "fmt"

func main() {
	sli1 := make([]int, 5, 10)
	for i := 0; i < len(sli1); i++ {
		sli1[i] = 2*i + 2
	}
	fmt.Println("slic1的元素为: ", sli1, "容量为: ", cap(sli1))
	fmt.Printf("%p\n", &(sli1[0]))

	sli3 := append(sli1, 1, 2, 3, 4, 5)
	fmt.Println("slic3的元素为: ", sli3, "容量为: ", cap(sli3))
	fmt.Printf("%p\n", &(sli3[0]))

	sli1 = append(sli1, sli1...)
	fmt.Println("slic1的元素为: ", sli1, "容量为: ", cap(sli1))
	fmt.Printf("%p\n", &(sli1[0]))

	sli2 := make([]int, 5, 10)
	fmt.Println("slic2的元素为: ", sli2, "容量为: ", cap(sli2))
	fmt.Printf("%p \n", &sli2)

	sli1 = append(sli1, sli1...)
	fmt.Println("slic1的元素为: ", sli1, "容量为: ", cap(sli1))
	fmt.Printf("%p \n", &(sli1[19]))

	slic3 := sli1[1:3:5]
	fmt.Printf("slic3的元素为: %v 长度为: %v 容量为: %v", slic3, len(slic3), cap(slic3))

}
