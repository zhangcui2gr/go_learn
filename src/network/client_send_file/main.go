package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func sendFile(file_name string, conn net.Conn) {
	f, err := os.Open(file_name)
	if err != nil {
		fmt.Println("err os.Open ", err)
		return
	}
	defer f.Close()
	buf := make([]byte, 1024*4)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("err os.Open ", err)
			}
			return
		}
		conn.Write(buf[:n])
	}

}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err net.Listen ", err)
		return
	}
	defer conn.Close()

	fmt.Println("plz input file name(include location)")
	var file_name string
	fmt.Scanln(&file_name)
	_, err1 := conn.Write([]byte(file_name))
	if err1 != nil {
		fmt.Println("err1 net.Listen ", err1)
		return
	}

	//等待接受服务器的"ok"字段
	buf := make([]byte, 1024)
	n, err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println("err2 conn.Read ", err2)
		return
	}
	if string(buf[:n]) == "ok" {
		fmt.Println("开始发送文件 ", file_name)
	}

	sendFile(file_name, conn)

}
