package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

/*
练习 1.10： 找一个数据量比较大的网站，用本小节中的程序调研网站的缓存策略，对每个URL执
行两遍请求，查看两次时间是否有较大的差别，并且每次获取到的响应内容是否一致，修改本节中
的程序，将响应结果输出，以便于进行对比。
*/

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("读取 %s 时出错: %v", url, err)
		return
	}

	defer resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
	nbytes = int64(len(body))
	secs = time.Since(start).Seconds()
}
