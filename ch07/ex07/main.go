package main

import (
	"flag"

	"fmt"

	"github.com/tokiw/golang-practice/ch07/ex07/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

// デフォルト値のときのCelsiusのString()を呼んでるので
// ここではCelsius.String()を書き換えた
func main() {
	flag.Parse()
	fmt.Println(*temp)
}
