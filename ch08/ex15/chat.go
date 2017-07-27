package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

//!+broadcaster
type client chan<- string // an outgoing message channel
const timeout = 5 * time.Second

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				select {
				case cli <- msg:
				default: // skip
				}
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	input := bufio.NewScanner(conn)

	go clientWriter(conn, ch)

	var name string
	ch <- "You are: "
	if input.Scan() {
		name = input.Text()
	} else {
		conn.Close()
		return
	}
	messages <- name + " has arrived"
	entering <- ch

	timer := time.NewTimer(timeout)
	go func() {
		<-timer.C
		conn.Close()
	}()
	for input.Scan() {
		messages <- name + ": " + input.Text()
		timer.Reset(timeout)
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- name + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
