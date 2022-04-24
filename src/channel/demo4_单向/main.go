package main

func main() {

	ch1 := make(chan<- int)
	ch1 <- 666

	ch2 := make(<-chan int)
	<-ch2

	//双向channel转单向channel
	ch := make(chan string)
	var ch3 chan<- string = ch
	ch3 <- "hello"

	var ch4 <-chan string = ch
	<-ch4

	//单向不能转双向
}
