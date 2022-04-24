package main

import (
	"fmt"
	"strings"
)

func Addupper() func(n1 int) int {
	var n int = 10
	var name string = "hello"
	return func(x int) int {
		n = n + x
		name += string(" ")
		fmt.Println(name)
		return n
	}
}

func makeSuffix(suffix string) func(name string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {

	f := Addupper()
	fmt.Println(f(1))
	fmt.Println(f(2))

	g := makeSuffix(".jpg")
	fmt.Printf("name = %v", g("video"))

}
