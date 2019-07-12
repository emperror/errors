package errors

import (
	"testing"
)

func TestNewPlain(t *testing.T) {
	err := NewPlain("error")

	t.Parallel()

	t.Run("error_message", func(t *testing.T) {
		checkErrorMessage(t, err, "error")
	})

	t.Run("format", func(t *testing.T) {
		checkFormat(t, err, map[string][]string{
			"%s":  {"error"},
			"%q":  {`"error"`},
			"%v":  {"error"},
			"%+v": {"error"},
		})
	})
}

func TestNew(t *testing.T) {
	err := New("error")
	origErr := err.(*withStack).error

	t.Parallel()

	t.Run("error_message", func(t *testing.T) {
		checkErrorMessage(t, err, "error")
	})

	t.Run("unwrap", func(t *testing.T) {
		checkUnwrap(t, err, origErr)
	})

	t.Run("format", func(t *testing.T) {
		checkFormat(t, err, map[string][]string{
			"%s":  {"error"},
			"%q":  {`"error"`},
			"%v":  {"error"},
			"%+v": {"error", "emperror.dev/errors.TestNew\n\t.+/errors_new_test.go:27"},
		})
	})
}

func TestErrorf(t *testing.T) {
	err := Errorf("error: %s", "something went wrong")
	origErr := err.(*withStack).error

	t.Parallel()

	t.Run("error_message", func(t *testing.T) {
		checkErrorMessage(t, err, "error: something went wrong")
	})

	t.Run("unwrap", func(t *testing.T) {
		checkUnwrap(t, err, origErr)
	})

	t.Run("format", func(t *testing.T) {
		checkFormat(t, err, map[string][]string{
			"%s":  {"error: something went wrong"},
			"%q":  {`"error: something went wrong"`},
			"%v":  {"error: something went wrong"},
			"%+v": {"error: something went wrong", "emperror.dev/errors.TestErrorf\n\t.+/errors_new_test.go:51"},
		})
	})
}
