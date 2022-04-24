package main

import "fmt"

var ch1 chan string = make(chan string)

func main() {
	////ch := make(chan int)
	//ch2 := make(chan string)
	//
	//
	////for str := range ch2 {
	////	fmt.Println(str)
	////}
	////<-ch2
	//go func() {
	//	num := <-ch2
	//	fmt.Println(num)
	//	fmt.Println("2")
	//}()
	//ch2 <- "string"
	//go func() {
	//	ch2 <- "11"
	//}()
	//fmt.Println("1")

	//go func() {
	//	defer fmt.Println("子协程结束")
	//	fmt.Println("子协程开始")
	//	ch <- 666
	//}()
	//num := <-ch
	//fmt.Println(num)
	//for {
	//
	//}

	go func() {
		num := <-ch1
		fmt.Println(num)
	}()
	ch1 <- "string"
	fmt.Println("执行主协程")

}
