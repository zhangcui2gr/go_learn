package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8000")
	if err != nil {
		fmt.Println("http.Get err ", err)
	}
	defer resp.Body.Close()

	fmt.Println("header = ", resp.Header)
	fmt.Println("status = ", resp.Status)
	fmt.Println("statuscode = ", resp.StatusCode)
	fmt.Printf("body type = %T\n", resp.Body)

	buf := make([]byte, 1024)
	var str string
	for {
		n, err2 := resp.Body.Read(buf)
		if err2 != nil {
			fmt.Println("resp.Body.Read ", err2)
		}
		if n == 0 {
			fmt.Println("End of Read")
			break
		}
		str += string(buf[:n])
	}

	fmt.Println("str ", str)

}
