package pio_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/tada/catch"
	"github.com/tada/catch/pio"
)

func TestWriteBool(t *testing.T) {
	w := bytes.Buffer{}
	pio.WriteBool(true, &w)
	pio.WriteBool(false, &w)
	a := w.String()
	if a != "truefalse" {
		t.Error("expected 'truefalse', got", a)
	}
}

func TestWriteInt(t *testing.T) {
	w := bytes.Buffer{}
	pio.WriteInt(12345, &w)
	a := w.String()
	if a != "12345" {
		t.Error("expected '12345', got", a)
	}
}

func TestWriteFloat(t *testing.T) {
	w := bytes.Buffer{}
	pio.WriteFloat(3.14159, &w)
	a := w.String()
	if a != "3.14159" {
		t.Error("expected '3.14159', got", a)
	}
}

func TestWriteRune(t *testing.T) {
	w := bytes.Buffer{}
	pio.WriteRune('a', &w)
	pio.WriteRune('⌘', &w)
	pio.WriteRune('x', &w)
	a := w.String()
	if a != "a⌘x" {
		t.Error("expected 'a⌘x', got", a)
	}
}

type badWriter int

func (b badWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New("bad write")
}

func TestWriteError(t *testing.T) {
	err := catch.Do(func() {
		pio.WriteRune('⌘', badWriter(0))
	})
	if err == nil {
		t.Fatal("expected error")
	}
}