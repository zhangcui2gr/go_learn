package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
			fmt.Println("已经写入了 ", i)
		}
		close(ch)
	}()
	for {
		for num := range ch {
			fmt.Println("num = ", num)
		}
		if val, ok := <-ch; ok == true {
			fmt.Println("val = ", val)
		} else {
			fmt.Println("溜了溜了")
			break
		}
	}

}
