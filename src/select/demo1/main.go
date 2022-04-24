package main

import "fmt"

//c通道负责传入的次数
//quit通道负责控制结束的时机
func feibo(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
			//fmt.Println(x, y)
		case a := <-quit:
			fmt.Println("退出 ", a)
			return
		}
	}
}

func main() {
	/*
		x, y := 1, 2
		//相当于python中的交换
		x, y = y, x+y+10
		fmt.Println(x, y)
	*/

	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 8; i++ {
			//fmt.Println(<-c)
			num := <-c
			fmt.Println(num)
		}
		quit <- 666
	}()

	feibo(c, quit)
}
