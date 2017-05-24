package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	var isSha256, isSha384, isSha512 bool
	flag.BoolVar(&isSha256, "sha256", false, "sha256")
	flag.BoolVar(&isSha384, "sha384", false, "sha384")
	flag.BoolVar(&isSha512, "sha512", false, "sha512")

	flag.Parse()

	var str string
	if len(flag.Args()) > 0 {
		str = flag.Args()[0]
	} else {
		fmt.Printf("Usage: -shaXXX {string} ")
		return
	}
	if isSha256 && !isSha384 && !isSha512 {
		fmt.Printf("%s: %x\n", str, sha256.Sum256([]byte(str)))
	} else if !isSha256 && isSha384 && !isSha512 {
		fmt.Printf("%s: %x\n", str, sha512.Sum384([]byte(str)))
	} else if !isSha256 && !isSha384 && isSha512 {
		fmt.Printf("%s: %x\n", str, sha512.Sum512([]byte(str)))
	} else {
		fmt.Printf("Usage: -shaXXX {string} ")
		return
	}
}
