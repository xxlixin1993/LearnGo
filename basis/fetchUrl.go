package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"time"
	"io"
)

func main() {
	getOne()
	getMore()
}

func getOne() {
	url := "http://www.baidu.com"
	response, err := http.Get(url)

	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}

	fmt.Println(string(body))
}

func getMore() {
	start := time.Now()
	uri := []string{
		"http://www.baidu.com",
		"http://www.taobao.com",
	}
	ch := make(chan string)

	for _, url := range uri {
		go fetch(url, ch)
	}

	for range uri {
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

	// 将resp.Body复制到ioutil.Discard（即>/dev/null）
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
