package main

import (
	"fmt"

)

var quit = make(chan int)

func main() {
	go t()
	go t()

	<-quit
	<-quit
}

func t() {
	for i := 0; i < 100; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	quit <- 1
}
