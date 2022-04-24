package main

import (
	"fmt"
	"time"
)

func main() {

	time1 := time.Now()
	fmt.Printf("time = %v, type = %T \n", time1, time1)

	fmt.Printf("Unix = %v\n", int(time1.Weekday()))

	fmt.Println(time1.Unix())

}
