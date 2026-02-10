package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rr.r.Read(b)
	for i := 0; i < len(b); i++ {
		b[i] = rot13(b[i])
	}
	return n, err
}

func rot13(c byte) byte {
	switch {
	case 'a' <= c && c <= 'z':
		return 'a' + (c-'a'+13)%26
	case 'A' <= c && c <= 'A':
		return 'A' + (c-'A'+13)%26
	default:
		return c
	}
}

func main() {
	s := strings.NewReader("xbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
