package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(2 * time.Second)
	curTime := time.Now()
	fmt.Println(curTime)

	t := <-timer.C
	fmt.Println(t)

}
