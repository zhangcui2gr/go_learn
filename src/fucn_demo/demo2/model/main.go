package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generate_randn(seed int64, num int) int {
	rand.Seed(seed)
	return rand.Intn(num)
}

func test1(n int) {
	if n > 2 {
		n--
		test1(n)
	}
	fmt.Println("n = ", n)

}

func test2(n int) {
	if n > 2 {
		n--
		test1(n)
	} else {
		fmt.Println("n = ", n)
	}

}

func func_feibo(n int) int {

	if n <= 2 {
		return 1
	}
	a, b := 1, 1
	for i := 2; i < n; i++ {
		c := a + b
		a = b
		b = c
	}

	return b
}

func func_test1(n int) int {
	if n == 1 {
		return 3
	}
	a := 3
	for i := 1; i < n; i++ {
		a = 2*a + 1
	}
	return a

}

func func_test2(args ...int) int {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func func_test3(n1, n2 float32) float32 {
	fmt.Printf("n1 type = %T\n", n1)
	return n1 + n2
}

type mySum func(int, int) int

func sum(n1, n2 int) int {
	return n1 + n2
}

func sum2(n1, n2 int) int {
	return n1 + n2
}

func myFunc(funcVar mySum, n1 int, n2 int) int {
	return funcVar(n1, n2)
}

func mySwap(n1, n2 *int) (int, int) {
	t := *n1
	*n1 = *n2
	*n2 = t
	return *n1, *n2
}

func main() {

	rand_num := generate_randn(time.Now().UnixMicro(), 100)
	fmt.Println(rand_num)

	test2(4)
	fmt.Println(func_feibo(7))

	fmt.Println(func_test1(2))

	fmt.Println(func_test2(2, 3, 4))

	fmt.Println(func_test3(1, 2))

	a := sum
	b := sum2

	fmt.Println(myFunc(a, 1, 2))
	fmt.Println(myFunc(b, 1, 2))

	num1, num2 := 1, 2

	fmt.Println(mySwap(&num1, &num2))

}
