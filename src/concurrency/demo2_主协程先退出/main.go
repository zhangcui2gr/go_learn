package main

import (
	"fmt"
	"time"
)

func Test() {
	for {
		fmt.Println("正在调用Test函数")
		time.Sleep(time.Second)
	}
}

func main() {

	go Test()
	go func() {
		fmt.Println("hello")
	}()

	fmt.Println("exit!")

}
