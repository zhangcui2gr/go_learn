package main

import (
	"fmt"
)

func producer(ch chan<- int) {
	defer close(ch)
	for i := 0; i < 5; i++ {
		ch <- i
		println("已经生产 ", i)
	}
}

func consumer(ch <-chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}

func main() {

	ch := make(chan int)
	go producer(ch)
	consumer(ch)

	fmt.Println("done")

}
