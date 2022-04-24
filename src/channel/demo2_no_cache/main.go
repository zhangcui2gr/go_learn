package main

import (
	"fmt"
	"time"
)

func main() {
	//有缓存channel
	ch := make(chan int, 3)

	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))

	go func() {
		for i := 0; i <= 10; i++ {
			ch <- i
			fmt.Println("i = ", i)
			fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))
		}
	}()

	time.Sleep(time.Second * 2)

	for i := 0; i <= 5; i++ {
		num := <-ch
		fmt.Println(num)
	}
}
