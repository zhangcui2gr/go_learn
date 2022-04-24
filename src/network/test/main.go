package main

import "fmt"

func main() {
	for {
		var str string
		_, err := fmt.Scanln(&str)
		if err != nil {
			fmt.Println("err fmt.Scanln = ", err)
			return
		}
		fmt.Println("input ", str)
	}

}
