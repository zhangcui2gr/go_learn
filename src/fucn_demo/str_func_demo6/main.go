package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var num int
	var str1 = "1234在"
	r := []rune(str1)

	num, _ = strconv.Atoi("12")

	for i := 0; i < len(str1); i++ {
		fmt.Printf("%c ", str1[i])
	}

	fmt.Println(" ")

	for i := 0; i < len(r); i++ {
		fmt.Printf("%c ", r[i])
	}

	fmt.Println(" ")

	var byte1 = []byte("hello go")
	fmt.Printf("byte1 = %s \n", byte1)

	var str2 = string([]byte{97, 98, 99})
	fmt.Println("str2 = ", str2)

	var num1 = strconv.FormatInt(123, 16)
	fmt.Println("num1 = ", num1)

	var b1 = strings.Contains("seafood", "foo")
	fmt.Println("num2 = ", b1)

	var num3 = strings.Count("hello", "l")
	fmt.Println("num3 = ", num3)

	var b2 = strings.EqualFold("abc", "AbC")
	fmt.Println("b2 = ", b2)

	var str3 = strings.TrimSpace(" hello world ")
	fmt.Println("str3 = ", str3)

	var str4 = strings.Trim(" !hello world!", "!")
	fmt.Println("str4 = ", str4)

	var str5 = strings.TrimLeft("!hello world!", "!")
	fmt.Println("str5 = ", str5)

	var str6 []string = strings.Split("hello,world,yes", ",")
	fmt.Printf("str6 = %T \n", str6)

	fmt.Println(num)
	fmt.Println(str1, " 的len长度是: ", len(str1))
	fmt.Println(r, " 的len长度是: ", len(r))

}
