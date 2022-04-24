package main

import (
	"fmt"
	"os"
)

func main() {

	list0 := os.Args
	if len(list0) != 2 {
		fmt.Println("plz input file name")
		return
	}
	file_name := list0[1]
	info, err := os.Stat(file_name)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("name = ", info.Name())
	fmt.Println("mode = ", info.Mode())
	fmt.Println("isDir = ", info.IsDir())

}
