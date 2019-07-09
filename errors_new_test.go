package errors

import (
	"testing"
)

func TestNewPlain(t *testing.T) {
	err := NewPlain("msg")

	if got, want := err.Error(), "msg"; got != want {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

func TestNew(t *testing.T) {
	err := New("msg")
	origErr := err.(*withStack).error

	if got, want := err.Error(), "msg"; got != want {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}

	testUnwrap(t, err, origErr)
	testFormatRegexp(t, 1, err, "%+v", "msg\nemperror.dev/errors.TestNew\n\t.+/errors_new_test.go:16")
}

func TestErrorf(t *testing.T) {
	err := Errorf("msg: %s", "msg2")
	origErr := err.(*withStack).error

	if got, want := err.Error(), "msg: msg2"; got != want {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}

	testUnwrap(t, err, origErr)
	testFormatRegexp(t, 1, err, "%+v", "msg\nemperror.dev/errors.TestErrorf\n\t.+/errors_new_test.go:28")
}
