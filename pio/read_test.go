package pio_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/tada/catch"

	"github.com/tada/catch/pio"
)

func TestRead(t *testing.T) {
	r := bytes.NewReader([]byte{1, 2, 3, 4})
	x := make([]byte, 3)
	n := pio.Read(x, r)
	if n != 3 {
		t.Fatalf("expected Read() to return 3, got %d", n)
	}
	if !bytes.Equal([]byte{1, 2, 3}, x) {
		t.Fatalf("expected Read() result to be []{1, 2, 3}, got %v", x)
	}
	n = pio.Read(x, r)
	if n != 1 {
		t.Fatalf("expected Read() to return 1, got %d", n)
	}
	if !bytes.Equal([]byte{4}, x[:1]) {
		t.Fatalf("expected Read() result to be []{4}, got %v", x[:1])
	}
	n = pio.Read(x, r)
	if n != -1 {
		t.Fatalf("expected Read() to return -1, got %d", n)
	}
}

type badReader int

func (b badReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("bad read")
}

func TestRead_badReader(t *testing.T) {
	r := badReader(0)
	x := make([]byte, 3)
	err := catch.Do(func() {
		pio.Read(x, r)
	})
	if err == nil || err.Error() != "bad read" {
		t.Fatalf(`expected Read() to error with "bad read", got %v`, err)
	}
}

func TestReadByte(t *testing.T) {
	r := bytes.NewReader([]byte{1, 2})
	b := pio.ReadByte(r)
	if b != 1 {
		t.Fatalf("expected ReadByte() to return 1, got %d", b)
	}
	b = pio.ReadByte(r)
	if b != 2 {
		t.Fatalf("expected ReadByte() to return 2, got %d", b)
	}
	b = pio.ReadByte(r)
	if b != -1 {
		t.Fatalf("expected ReadByte() to return -1, got %d", b)
	}
}

func TestReadByte_badReader(t *testing.T) {
	r := badReader(0)
	err := catch.Do(func() {
		pio.ReadByte(r)
	})
	if err == nil || err.Error() != "bad read" {
		t.Fatalf(`expected ReadByte() to error with "bad read", got %v`, err)
	}
}
