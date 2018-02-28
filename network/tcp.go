package main

import (
	"github.com/xxlixin1993/CacheGo/logging"
	"net"
	"fmt"
)

func main() {
	StartTcpServer()
}

func StartTcpServer() error {
	var listenErr error
	tcpListener, listenErr := net.Listen("tcp", "127.0.0.1:2222")

	if listenErr != nil {
		return listenErr
	}

	defer tcpListener.Close()

	for {
		conn, acceptErr := tcpListener.Accept()
		if acceptErr != nil {
			logging.Error("[server] Accept error, msg: ", acceptErr)
		}
		go connHandler(conn)
	}

	return nil
}

func connHandler(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			//break
			return
		}
		fmt.Printf("Received data: %v", string(buf[:len]))
	}
}
