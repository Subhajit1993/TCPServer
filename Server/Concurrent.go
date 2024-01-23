package Server

import (
	"fmt"
	"net"
	"runtime"
	"time"
)

func handleConcurrentConnection(conn net.Conn) {
	defer conn.Close()
	time.Sleep(2 * time.Second)
	fmt.Println("Connection handled")
}

func (s Concurrent) Start() (bool, error) {
	listener, err := net.Listen("tcp", ":"+s.config.Port)
	if err != nil {
		fmt.Println("Error:", err)
		return false, err
	}
	defer listener.Close()

	fmt.Println("Server listening on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Println("Number of active goroutines:", runtime.NumGoroutine())
		go handleConcurrentConnection(conn)
	}
}
