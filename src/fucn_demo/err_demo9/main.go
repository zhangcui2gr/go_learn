package main

import (
	"errors"
	"fmt"
)

func test() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err = ", err)
		}
	}()

	num1, num2 := 10, 0
	num3 := num1 / num2
	fmt.Println("num3 = ", num3)
}

func resd_file(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误")
	}
}

func test02() {
	// defer func() {
	// 	err := recover()
	// 	if err != nil {
	// 		fmt.Println("err = ", err)
	// 	}
	// }()

	err := resd_file("config")
	if err != nil {
		panic(err)
	}
	fmt.Println("test02()执行到这了")

}

func main() {

	test()
	test02()
	fmt.Println("执行到这了")

}
