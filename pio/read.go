package pio

import (
	"io"

	"github.com/tada/catch"
)

// ReadByte reads and returns the next byte from the input as an int ranging from 0 to 255. -1 is returned
// when the reader reaches EOF and no byte was read.
//
// If an error occurs the method panics with a Error with the Cause set to that error
func ReadByte(r io.Reader) int {
	b := []byte{0}
	n, err := r.Read(b)
	if err == nil || err == io.EOF {
		if n == 0 {
			return -1
		}
		return int(b[0])
	}
	panic(catch.Error{Cause: err})
}

// Read reads up to len(p) bytes into p. It returns the number of bytes read (0 <= n <= len(p)). -1 is returned when
// the reader reaches EOF and the number of bytes read is zero.
//
// If an error occurs the method panics with a Error with the Cause set to that error
func Read(p []byte, r io.Reader) int {
	n, err := r.Read(p)
	if err != nil {
		if err != io.EOF {
			panic(catch.Error{Cause: err})
		}
		if n == 0 {
			n = -1
		}
	}
	return n
}
