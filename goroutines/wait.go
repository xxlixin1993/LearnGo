package main

import (
	"sync"
	"fmt"
)

type message struct {
	body chan []byte
	sync.WaitGroup
}

var msg = &message{
	body: make(chan []byte, 1000),
}

func main() {
	msg.Add(1)
	go run()
	msg.body <- []byte("adsadasdas")
	close(msg.body)

	// 等待协程结束 再退出
	msg.Wait()
}

func run() {
	for {
		body, ok := <-msg.body
		if !ok {
			msg.Done()
			break
		}
		fmt.Println(body)
	}
}
