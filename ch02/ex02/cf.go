package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/tokiw/golang-practice/ch02/ex01/tempconv"
	"github.com/tokiw/golang-practice/ch02/ex02/lengthconv"
	"github.com/tokiw/golang-practice/ch02/ex02/weigthconv"
)

func main() {
	input := os.Args[1:]
	if len(input) == 0 {
		fmt.Print("value > ")
		var in string
		fmt.Scan(&in)
		input = make([]string, 1)
		input[0] = in
	}
	for _, arg := range input {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		showTemp(t)
		showLength(t)
		showWeigth(t)
	}
}

func showTemp(temp float64) {
	f := tempconv.Fahrenheit(temp)
	c := tempconv.Celsius(temp)
	fmt.Printf("%s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c))
}

func showLength(len float64) {
	f := lengthconv.Foot(len)
	m := lengthconv.Meter(len)
	fmt.Printf("%s = %s, %s = %s\n",
		f, lengthconv.FToM(f), m, lengthconv.MToF(m))
}

func showWeigth(weigth float64) {
	k := weigthconv.Kilogram(weigth)
	p := weigthconv.Pound(weigth)
	fmt.Printf("%s = %s, %s = %s\n",
		k, weigthconv.KToP(k), p, weigthconv.PToK(p))
}
