package main

import (
	"fmt"
	"net"
	"os"
)

func recvFile(file_name string, conn net.Conn) {
	newFile, err := os.Create(file_name)
	defer newFile.Close()
	if err != nil {
		fmt.Println("err os.Create ", err)
		return
	}
	buf := make([]byte, 1024*4)
	for {
		n, err1 := conn.Read(buf)
		if err1 != nil {
			fmt.Println("err1 conn.Read ", err1)
			return
		}
		if n == 0 {
			fmt.Println("n == 0")
			return
		}
		newFile.Write(buf[:n])
	}

}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err net.Listen ", err)
		return
	}
	defer listener.Close()

	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("err1 listener.Accept ", err1)
		return
	}
	defer conn.Close()
	buf := make([]byte, 1024*4)
	n, err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println("err2 conn.Read ", err2)
		return
	}
	info, err3 := os.Stat(string(buf[:n]))
	if err3 != nil {
		fmt.Println("err3 os.Stat ", err3)
		return
	}
	file_name := info.Name()
	conn.Write([]byte("ok"))
	fmt.Println("开始接收文件 ", file_name)

	recvFile(string(file_name), conn)

}
