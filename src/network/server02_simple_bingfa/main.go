package main

import (
	"fmt"
	"net"
	"strings"
)

func Hanndle(conn net.Conn) {
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, " connect successful")
	defer conn.Close()

	go func() {
		for {
			str := ""
			fmt.Scanln(&str)
			_, err := conn.Write([]byte(str))
			if err != nil {
				fmt.Println("err = ", err)
				return
			}
		}
	}()

	for {
		buf := make([]byte, 2048)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		fmt.Printf("%s 发送 %s\n", addr, string(buf[:n]))
		//fmt.Println(len(string(buf[:n])))
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
		if string(buf[:n]) == "exit" {
			fmt.Println(addr, " exit successful")
			conn.Close()
			return
		}
	}
}

func main() {
	listener, err1 := net.Listen("tcp", "127.0.0.1:8000")
	if err1 != nil {
		fmt.Println("err = ", err1)
		return
	}
	defer listener.Close()

	for {
		conn, err2 := listener.Accept()
		if err2 != nil {
			fmt.Println("err2 = ", err2)
			return
		}
		go Hanndle(conn)
	}

}
