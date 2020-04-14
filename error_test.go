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
