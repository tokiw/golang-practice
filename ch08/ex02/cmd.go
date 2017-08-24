package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func userCmd(op []string, s *status) (string, int, error) {
	s.user = op[0]

	if op[0] == "anonymous" || op[0] == "ftp" {
		return "Guest login ok, type your name as password.", 331, nil
	}
	return "User " + op[0] + " accepted, provide password.", 331, nil
}

func passCmd(op []string, s *status) (string, int, error) {
	if s.user == "anonymous" || s.user == "ftp" {
		return "anonymous don't require pass.", 230, nil
	}

	return "Login incorrect", 530, nil
}

func portCmd(op []string, s *status) (string, int, error) {
	ops := strings.Split(op[0], ",")

	p1, _ := strconv.Atoi(ops[4])
	p2, _ := strconv.Atoi(ops[5])
	s.host = fmt.Sprintf("%s.%s.%s.%s", ops[0], ops[1], ops[2], ops[3])
	s.port = p1*256 + p2

	return fmt.Sprintf("PORT command successful"), 200, nil
}

func quitCmd(op []string, s *status) (string, int, error) {
	return "Thank you for using the FTP service", 221, nil
}

func retrCmd(op []string, s *status) (string, int, error) {
	if len(op) < 1 {
		return "RETR command is required a parammeter.", 500, nil
	}

	addr := fmt.Sprintf("%s:%d", s.host, s.port)

	path := filepath.Join(s.path, op[0])

	file, err := os.Open(path)

	defer file.Close()
	if err != nil {
		return "", 550, err
	}

	s.resp <- response{"File ok", 150, nil}

	conn, err := net.Dial("tcp", addr)
	defer conn.Close()
	if err != nil {
		return "", 425, err
	}

	io.Copy(conn, file)
	return "Success receive file", 226, nil
}

func cwdCmd(op []string, s *status) (string, int, error) {
	path := filepath.Join(s.path, op[0])

	if f, err := os.Stat(path); err != nil || !f.IsDir() {
		return "Failed dir changed", 550, err
	}

	s.path = path
	return "PORT command successful", 250, nil
}

func listCmd(op []string, s *status) (string, int, error) {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)

	path := s.path
	if len(op) > 0 {
		path = filepath.Join(path, op[0])
	}

	stat, err := os.Stat(path)
	if err != nil {
		return s.path + " cant open.", 550, err
	}
	var str string
	if stat.IsDir() {
		dir, err := ioutil.ReadDir(path)
		if err != nil {
			return path + "cant read.", 550, err
		}

		for _, d := range dir {
			str += d.Name() + "\n"
		}
	} else {
		str = path
	}
	s.resp <- response{"Opening data connection.", 150, nil}

	conn, err := net.Dial("tcp", addr)
	defer conn.Close()
	if err != nil {
		return "", 425, err
	}

	_, err = fmt.Fprintf(conn, "%s\n", str)
	if err != nil {
		return path + "cant send.", 550, err
	}

	return "Transfer complete", 226, nil
}
