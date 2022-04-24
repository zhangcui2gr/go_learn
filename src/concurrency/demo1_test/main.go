package main

import (
	"fmt"
	"time"
)

func Newtask() {
	for {
		fmt.Println("正在执行Newtask函数")
		time.Sleep(time.Second)
	}
}

func main() {

	go Newtask()
	for i := 0; i < 10; i++ {
		fmt.Printf("正在第%d次执行main函数\n", i+1)
		time.Sleep(time.Second)
	}

}
