package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCharCount(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"a", "rune\tcount\n'a'\t1\n\nlen\tcount\n1\t1\n2\t0\n3\t0\n4\t0\n"},
		{"aba", "rune\tcount\n'a'\t2\n'b'\t1\n\nlen\tcount\n1\t3\n2\t0\n3\t0\n4\t0\n"},
		{"あ", "rune\tcount\n'あ'\t1\n\nlen\tcount\n1\t0\n2\t0\n3\t1\n4\t0\n"},
	}

	for _, test := range tests {
		stdout, stderr := StubIO(test.input, func() {
			main()
		})
		if len(stderr) != 0 {
			t.Errorf("error: %q", stderr)
		}
		if got := stdout; got != test.want {
			t.Errorf("input: %q\nresult: %q, expected: %q", test.input, got, test.want)
		}
	}
}

func StubIO(input string, fn func()) (string, string) {
	inr, inw, _ := os.Pipe()
	outr, outw, _ := os.Pipe()
	errr, errw, _ := os.Pipe()
	// 退避
	orgStdin := os.Stdin
	orgStdout := os.Stdout
	orgStderr := os.Stderr
	inw.Write([]byte(input))
	inw.Close()
	// 上書き
	os.Stdin = inr
	os.Stdout = outw
	os.Stderr = errw
	fn()
	// 後処理
	os.Stdin = orgStdin
	os.Stdout = orgStdout
	os.Stderr = orgStderr
	outw.Close()
	outbuf, _ := ioutil.ReadAll(outr)
	errw.Close()
	errbuf, _ := ioutil.ReadAll(errr)

	return string(outbuf), string(errbuf)
}
