package main

import (
	"fmt"
	"net"
)

type Client struct {
	C    chan string
	Name string
	Addr string
}

var OnlineMap map[string]Client

var message = make(chan string)

func WriteToClient(cli Client, conn net.Conn) {
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMessage(cli Client, str string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ":" + str
	return buf
}

func Handdle(conn net.Conn) {
	defer conn.Close()
	cliAddr := conn.RemoteAddr().String()
	cli := Client{make(chan string), cliAddr, cliAddr}
	OnlineMap[cliAddr] = cli

	go WriteToClient(cli, conn)

	message <- MakeMessage(cli, "上线了")

	cli.C <- MakeMessage(cli, "i am this")

	//读到的消息写入message
	go func() {
		buf := make([]byte, 1024)

		for {
			n, err := conn.Read(buf)
			if n == 0 {
				fmt.Println("conn.Read err = ", err)
				return
			}
			msg := string(buf[:n])
			fmt.Println(msg)
			if msg == "who" && len(msg) == 3 {
				user := "online user : "
				for _, cli := range OnlineMap {
					user += "[" + cli.Addr + "]" + cli.Name + "   "
				}
				conn.Write([]byte(user + "\n"))
			} else {
				message <- MakeMessage(cli, msg)
			}
		}
	}()

	for {

	}
}

func Messager() {
	OnlineMap = make(map[string]Client)
	for {
		msg := <-message
		for _, cli := range OnlineMap {
			cli.C <- msg

		}
	}
}

func ReadMessage(conn net.Conn) []byte {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if n == 0 {
		fmt.Println("conn.Read err = ", err)
		return []byte("no")
	}
	msg := buf[:n]
	return msg
}

func main() {
	listener, err01 := net.Listen("tcp", "127.0.0.1:8000")
	if err01 != nil {
		fmt.Println("err01 net.Listen = ", err01)
		return
	}
	go Messager()
	defer listener.Close()
	for {
		conn, err02 := listener.Accept()
		if err02 != nil {
			fmt.Println("err02 listener.Accept = ", err02)
			continue
		}
		go Handdle(conn)

	}

}
