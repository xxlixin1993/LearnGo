package main

import (
	"net/http"
	"fmt"
	"time"
	"io"
	"os"
	"os/signal"
	"context"
)
var server *http.Server

func main() {
	go func() {
		server = &http.Server{
			Addr:         "127.0.0.1:4321",
			Handler:      getServerMux(),
			ReadTimeout:  time.Duration(3) * time.Second,
			WriteTimeout: time.Duration(3) * time.Second,
		}
		server.ListenAndServe()
	}()

	waitSignal()
}

func getServerMux() *http.ServeMux {
	mx := http.NewServeMux()
	mx.Handle("/", &LHandle{})
	return mx
}

type LHandle struct {
}

func (L *LHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Proto)
	io.WriteString(w,"hello")
}

// Wait signal
func waitSignal() {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan)

	sig := <-sigChan

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(2)*time.Second)

	server.Shutdown(ctx)
	fmt.Println("stop...")
	fmt.Printf("signal is %d", sig)
}
