package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(3)

	go func(s1 string) {
		for i := 0; i < 3; i++ {
			//fmt.Printf("正在执行%d次func", i)
			defer wg.Done()
			fmt.Println(s1)
			//time.Sleep(time.Second)
		}

	}("world")

	for i := 0; i < 3; i++ {
		//runtime.Gosched()
		fmt.Println("hello")
		wg.Wait()
		//time.Sleep(time.Second)
	}
}
