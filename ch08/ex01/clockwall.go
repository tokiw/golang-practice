package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	done := make(chan struct{})
	for i, arg := range os.Args[1:] {
		go startConn(arg, done, i)
	}

	for range os.Args[1:] {
		<-done
	}
}

func startConn(server string, done chan<- struct{}, number int) {
	fmt.Printf("Start: No.%d %s\n", number, server)
	defer func() {
		done <- struct{}{}
	}()

	conn, err := net.Dial("tcp", server)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
