package main

import (
	"golang.org/x/tour/reader"
	"strings"
)

type MyReader struct{}

func (reader MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
