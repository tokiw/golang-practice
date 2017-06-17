package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	str1 := "aa bb"
	fmt.Println(string(replaceSpace([]byte(str1))))
	str2 := "あ　い"
	fmt.Println(string(replaceSpace([]byte(str2))))
}

// 別のスペースを返しているのでだめ！！
func replaceSpace(s []byte) []byte {
	runes := bytes.Runes(s)

	for i := 0; i < len(runes); i++ {
		if unicode.IsSpace(runes[i]) {
			runes[i] = ' '
		}
	}
	return []byte(string(runes))
}
