package catch_test

import (
	"errors"
	"testing"

	"github.com/tada/catch"
)

func TestError(t *testing.T) {
	err := errors.New("the error")
	ex := catch.Error(err)
	if err.Error() != ex.Error() {
		t.Error("Error() created error produces wrong error string")
	}
	if !catch.IsError(ex) {
		t.Error("IsError() returns false for Error() created error ")
	}
	if catch.Cause(ex) != err {
		t.Error("Cause() does not return original error for Error() created error ")
	}
	if catch.Cause(err) != nil {
		t.Error("Cause() does not return nil on unknown errors")
	}
}

func TestError_string(t *testing.T) {
	ex := catch.Error("the error")
	if ex.Error() != "the error" {
		t.Error("Error() created error produces wrong error string")
	}
	if !catch.IsError(ex) {
		t.Error("IsError() returns false for Error() created error ")
	}
}

func TestError_fmt(t *testing.T) {
	ex := catch.Error("the numbers %d and %d are illegal", 38, 70)
	if ex.Error() != "the numbers 38 and 70 are illegal" {
		t.Error("Error() created error produces wrong error string")
	}
}

func TestError_noArgs(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil || err.(error).Error() != "catch.Error(): illegal argument" {
			t.Errorf("expected illegal argument panic")
		}
	}()
	_ = catch.Error()
}

func TestError_badArgs(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil || err.(error).Error() != "catch.Error(): illegal argument" {
			t.Errorf("expected illegal argument panic")
		}
	}()
	_ = catch.Error(errors.New("oops"), "bad additional arg")
}
