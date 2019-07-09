package errors

import (
	"testing"
)

type errorString struct {
	msg string
}

func (e *errorString) Error() string {
	return e.msg
}

func testUnwrap(t *testing.T, err error, origErr error) {
	t.Helper()

	if err, ok := err.(interface{ Unwrap() error }); ok {
		if got, want := err.Unwrap(), origErr; got != want {
			t.Errorf("error does not match the expected one\nactual:   %#v\nexpected: %#v", got, want)
		}
	} else {
		t.Fatal("error does not implement the wrapper (interface{ Unwrap() error}) interface")
	}

	if err, ok := err.(interface{ Cause() error }); ok {
		if got, want := err.Cause(), origErr; got != want {
			t.Errorf("error does not match the expected one\nactual:   %#v\nexpected: %#v", got, want)
		}
	} else {
		t.Fatal("error does not implement the causer (interface{ Cause() error}) interface")
	}
}
