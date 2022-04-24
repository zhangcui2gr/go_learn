package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func Printer(str string) {
	for _, val := range str {
		fmt.Printf("%c", val)
		time.Sleep(time.Microsecond * 500000)
	}
	fmt.Println()
}

func p1() {
	Printer("hello")
	ch <- 666
}

func p2() {
	//从管道读数据，没有则阻塞
	<-ch
	Printer("world")
}

func main() {

	//go Printer("hello")

	//go Printer("world")

	go p1()
	go p2()

	for {
	}

}
