package catch_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/tada/catch"
)

func ExampleDo() {
	err := catch.Do(func() {
		if err := os.Chmod("nosuchfile.noop", 0600); err != nil {
			panic(catch.Error{Cause: err})
		}
	})
	fmt.Println(err)
	// Output: chmod nosuchfile.noop: no such file or directory
}

func TestCatch_normal(t *testing.T) {
	e := errors.New("the error")
	err := catch.Do(func() {
		panic(&catch.Error{Cause: e})
	})
	if err != e {
		t.Errorf("expected %v, got %v", e, err)
	}
}

func TestCatch_normalByValue(t *testing.T) {
	e := errors.New("the error")
	err := catch.Do(func() {
		panic(catch.Error{Cause: e})
	})
	if err != e {
		t.Errorf("expected %v, got %v", e, err)
	}
}

func TestCatch_other(t *testing.T) {
	e := errors.New("the error")
	defer func() {
		err := recover()
		if err != e {
			t.Errorf("expected %v, got %v", e, err)
		}
	}()
	_ = catch.Do(func() {
		panic(e)
	})
}
