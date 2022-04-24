package main

import (
	"fmt"
	"net"
)

func WriteMessage(conn net.Conn) {
	for {
		str := ""
		fmt.Scanln(&str)
		conn.Write([]byte(str))
	}

}

func ReadMessage(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err1 := conn.Read(buf)
		if err1 != nil {
			fmt.Println("err1 conn.Read = ", err1)
			return
		}
		fmt.Println(string(buf[:n]))
	}

}

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err net.Dial = ", err)
		return
	}
	defer conn.Close()

	//子协程写
	go WriteMessage(conn)
	buf := make([]byte, 1024)
	//主协程读
	for {

		n, err1 := conn.Read(buf)
		if err1 != nil {
			fmt.Println("err1 conn.Read = ", err1)
			return
		}
		fmt.Println(string(buf[:n]))

	}
}
