package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

type CountWriter struct {
	writer io.Writer
	count  int64
}

func (b *CountWriter) Write(p []byte) (int, error) {
	c, err := b.writer.Write(p)
	b.count += int64(c)
	return c, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &CountWriter{}
	cw.writer = w
	return cw.writer, &cw.count
}

func main() {
	w, c := CountingWriter(ioutil.Discard)
	fmt.Println(*c)
	w.Write([]byte("hello"))
	fmt.Println(*c)
}
