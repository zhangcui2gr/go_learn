package main

import (
	"fmt"
	"net/http"
)

func myHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("okkk1")
	fmt.Println("client IP", req.RemoteAddr)
	fmt.Println("client Method: ", req.Method)
	//defer req.Body.Close()
	fmt.Println("client Body: ", req.Body)
	//defer req.Body.Close()
	fmt.Fprint(w, "Hello, HandlerFunc!")

	//log.Printf("Request received from %s", req.RemoteAddr)
	//w.Write([]byte("okkk"))
}

func main() {
	http.HandleFunc("/", myHandler)

	http.ListenAndServe("127.0.0.1:8000", nil)

}
