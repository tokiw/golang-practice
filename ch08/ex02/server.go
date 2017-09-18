package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type response struct {
	message string
	code    int
	err     error
}

type status struct {
	user     string
	host     string
	path     string
	port     int
	dataType string
	resp     chan response
}

type command func([]string, *status) (string, int, error)

var commands = map[string]command{
	"USER": userCmd,
	"PASS": passCmd,
	"QUIT": quitCmd,
	"PORT": portCmd,
	"RETR": retrCmd,
	"CWD":  cwdCmd,
	"LIST": listCmd,
}

// > ftp localhost 8000
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	var s status

	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(conn, "Fatal: %s", err.Error())
		return
	}
	s.path = dir

	input := bufio.NewScanner(conn)
	resp := make(chan response)
	s.resp = resp
	go writeResponse(conn, resp)

	fmt.Fprintf(conn, "220 Connected\n")

	for input.Scan() {
		line := input.Text()
		message, code, err := handleCmd(line, &s)
		if err != nil {
			fmt.Fprintf(conn, "%d %s\r\n", code, err.Error())
			continue
		}
		fmt.Fprintf(conn, "%d %s\r\n", code, message)
		if code == 221 {
			return
		}
	}
}s

func writeResponse(conn net.Conn, resp <-chan response) {
	for msg := range resp {
		fmt.Fprintf(conn, "%d %s\r\n", msg.code, msg.message)
	}
}

func handleCmd(line string, s *status) (message string, code int, err error) {
	name, op := parseCmd(line)
	cmd := commands[name]
	if cmd == nil {
		return fmt.Sprintf("%s is unsupported.", name), 500, nil
	}
	return cmd(op, s)
}

func parseCmd(line string) (string, []string) {
	lines := strings.Fields(line)
	return strings.ToUpper(lines[0]), lines[1:]
}
