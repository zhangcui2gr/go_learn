package main

import (
	"fmt"
	"net/http"
)

//服务端编写的业务逻辑处理程序
//hander函数： 具有func(w http.ResponseWriter, r *http.Requests)签名的函数
func myHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.RemoteAddr, "连接成功")  //r.RemoteAddr远程网络地址
	fmt.Println("method = ", r.Method) //请求方法
	fmt.Println("url = ", r.URL.Path)
	fmt.Println("header = ", r.Header)
	fmt.Println("body = ", r.Body)

	w.Write([]byte("hello go")) //给客户端回复数据
}

func main() {
	http.HandleFunc("/", myHandler)

	//该方法用于在指定的 TCP 网络地址 addr 进行监听，然后调用服务端处理程序来处理传入的连接请求。
	//该方法有两个参数：第一个参数 addr 即监听地址；第二个参数表示服务端处理程序，通常为空
	//第二个参数为空意味着服务端调用 http.DefaultServeMux 进行处理
	http.ListenAndServe("127.0.0.1:8000", nil)
}
