package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer conn.Close()
	for {

		go func() {
			var str string
			//fmt.Println("请输入需要发送的数据")
			fmt.Println("origin input ", str)
			//会出现读入数据和读到的数据不一致
			fmt.Scanf("%s\n", &str)
			//if err != nil {
			//	fmt.Println("err fmt.Scanln = ", err)
			//	return
			//}
			fmt.Println("input ", str)
			conn.Write([]byte(str))
		}()

		buf := make([]byte, 2048)
		n, err2 := conn.Read(buf)
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		fmt.Println(string(buf[:n]))
	}

}
