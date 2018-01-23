package main

import (
	"net/http"
	"fmt"
	"sync"
)

var mu sync.Mutex

var count int

func main() {
	http.HandleFunc("/count", viewCount)
	http.ListenAndServe("localhost:8000", nil)
}


func viewCount(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "Count %d\n", count)
}