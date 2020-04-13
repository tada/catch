package catch_test

import (
	"errors"
	"testing"

	"github.com/tada/catch"
)

func TestError(t *testing.T) {
	err := errors.New("the error")
	ex := catch.Error{Cause: err}
	if err.Error() != ex.Error() {
		t.Error("Error produces wrong error string")
	}
}
