package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	quit := make(chan bool)
	//通过select设置超时，因为只要满足其中一个就会执行
	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-time.After(time.Second * 3):
				fmt.Println("结束")
				quit <- true
				return
			}
		}
	}()

	c <- 5
	time.Sleep(time.Second * 2)
	c <- 6
	<-quit
}
