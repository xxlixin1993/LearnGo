package main

import (
	"net/http"
	"time"
	"fmt"
	"io"
)

func main() {
	startHttp()
}

func startHttp() {
	server := &http.Server{
		Addr:         "127.0.0.1:4321",
		Handler:      &ServerHandle{},
		ReadTimeout:  time.Duration(3) * time.Second,
		WriteTimeout: time.Duration(3) * time.Second,
	}
	server.ListenAndServe()
}

type ServerHandle struct {
}

func (sh *ServerHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Proto)
	io.WriteString(w, "hello")
}
