package main

import (
	"net/http"
	"fmt"
	"sync"
)
var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", testHandle)
	http.ListenAndServe("localhost:8000", nil)
}

func testHandle(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	count++
	fmt.Fprintf(w, "hello world, count is %d" , count)
	mu.Unlock()
}