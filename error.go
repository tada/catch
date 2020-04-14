package catch

// special error that is recovered by the Do() function returning its original cause.
type errorWithCause struct {
	// Cause is the original error
	cause error
}

// Error returns a new special error object that wraps the given cause. The returned error can be used as an
// argument to the IsError() and Cause() functions.
func Error(cause error) error {
	return &errorWithCause{cause: cause}
}

// Error returns the result of calling Error() on the contained cause
func (e *errorWithCause) Error() string {
	return e.cause.Error()
}

// IsError returns true if, and only if, the argument is an error produced by the Error function.
func IsError(e interface{}) bool {
	_, ok := e.(*errorWithCause)
	return ok
}

// Cause returns the underlying cause of the error provided that the argument is an error created by the Error function
// or nil if that is not the case.
func Cause(e interface{}) error {
	if e, ok := e.(*errorWithCause); ok {
		return e.cause
	}
	return nil
}
