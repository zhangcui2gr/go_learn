package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func httpGet(url string) (content string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http.Get err = ", err)
		return
	}
	defer resp.Body.Close()

	os.Chdir("D:\\career\\Golang\\code_demo\\src\\http")
	buf := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil {
			//fmt.Println("resp.Body.Read err = ", err)
			break
		}
		if n == 0 {
			//fmt.Println("读完了已经")
			break
		}
		content += string(buf)
	}
	return
}

func spiderPage(i, num int, AHU_tieba string, page chan int) {
	url := AHU_tieba + strconv.Itoa(i*num)
	fmt.Printf("正在爬取第%d个网页 : %s\n", i, url)
	content, err := httpGet(url)
	if err != nil {
		fmt.Println("httpGet err = ", err)
		return
	}
	file_name := strconv.Itoa(i) + ".html"
	f, err2 := os.Create(file_name)
	if err2 != nil {
		fmt.Println("os.Create err2 = ", err2)
		return
	}
	f.Write([]byte(content))
	f.Close()
	page <- i
}

func doWork(start_page, end_page int, AHU_tieba string) {
	num := 50
	page := make(chan int)
	for i := start_page - 1; i <= end_page-1; i++ {
		go spiderPage(i, num, AHU_tieba, page)
	}
	for i := start_page - 1; i <= end_page-1; i++ {
		fmt.Printf("第%d个页面爬取完毕\n", <-page)
	}
}

func main() {
	AHU_tieba := "https://tieba.baidu.com/f?kw=%E5%AE%89%E5%BE%BD%E5%A4%A7%E5%AD%A6&ie=utf-8&pn="
	fmt.Println("plz input start and end page")
	var start_page, end_page int
	fmt.Scanf("%d %d", &start_page, &end_page)
	fmt.Printf("开始 : %d 结束 : %d\n", start_page, end_page)
	doWork(start_page, end_page, AHU_tieba)

}
