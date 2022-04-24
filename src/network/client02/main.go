package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer conn.Close()
	for {
		str := ""
		fmt.Scanln(&str)
		conn.Write([]byte(str))
		buf := make([]byte, 2048)
		n, err2 := conn.Read(buf)
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		fmt.Println(string(buf[:n]))
	}

}
