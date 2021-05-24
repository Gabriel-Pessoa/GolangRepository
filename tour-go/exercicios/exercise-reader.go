package main

import (
	"fmt"
	"io"
)

type MyReader struct {
	R io.Reader
}

type ErrNegativeArray []byte

func (e ErrNegativeArray) Error() string {
	return fmt.Sprintf("cannot empty array: %v", len(e))
}

func (m MyReader) Read(b []byte) (int, error) {
	sizeInput := len(b)
	var n int = 0

	if sizeInput == 0 {
		return n, ErrNegativeArray(b)
	}

	for i := 0; i < sizeInput; i++ {
		fmt.Printf("n = %v err = %v b = %v\n", n, nil, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		n = i
	}
	return n, nil
}

func main() {
	reader.Validate(MyReader{})
}
