package pio

import (
	"io"
	"strconv"
	"unicode/utf8"

	"github.com/tada/catch"
)

// Write writes the bytes b to the Writer, returning its length.
//
// Any error from the io.Writer will result in a panic(catch.Error(err))
func Write(b []byte, w io.Writer) int {
	n, err := w.Write(b)
	if err != nil {
		panic(catch.Error(err))
	}
	return n
}

// WriteString writes the bytes of s to the Writer, returning its length.
//
// Any error from the io.Writer will result in a panic(catch.Error(err))
func WriteString(s string, w io.Writer) int {
	return Write([]byte(s), w)
}

// WriteByte writes the byte r to the Writer.
//
// Any error from the io.Writer will result in a panic(catch.Error(err))
func WriteByte(b byte, w io.Writer) {
	Write([]byte{b}, w)
}

// WriteRune writes the UTF-8 encoding of Unicode code point r to the Writer, returning its length.
//
// Any error from the io.Writer will result in a panic(catch.Error(err))
func WriteRune(r rune, w io.Writer) int {
	if r < utf8.RuneSelf {
		WriteByte(byte(r), w)
		return 1
	}
	b := make([]byte, utf8.UTFMax)
	n := utf8.EncodeRune(b, r)
	Write(b[:n], w)
	return n
}

// WriteBool writes the string "true" or "false" onto the stream
//
// Any error from the io.Writer will result in a panic(catch.Error(err))
func WriteBool(b bool, w io.Writer) {
	s := "false"
	if b {
		s = "true"
	}
	WriteString(s, w)
}

// WriteInt writes decimal string representation of the given integer onto the stream
//
// Any error from the io.Writer will result in a panic(catch.Error(err))
func WriteInt(i int64, w io.Writer) {
	WriteString(strconv.FormatInt(i, 10), w)
}

// WriteFloat writes the "%g" string representation of the given integer onto the stream
//
// Any error from the io.Writer will result in a panic(catch.Error(err))
func WriteFloat(f float64, w io.Writer) {
	WriteString(strconv.FormatFloat(f, 'g', -1, 64), w)
}
