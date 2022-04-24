package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second)
	timer.Reset(5 * time.Second)

	tik := time.NewTicker(1 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("时间到，子协程打印下")
		timer.Stop()
	}()

	for {
		go func() {
			<-tik.C
			fmt.Println("周期到，子协程打印下")
		}()

	}
}
