package main

import (
	"flag"

	"fmt"

	"github.com/tokiw/golang-practice/ch07/ex06/tempconv"
)

var temp = tempconv.KelvinFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
