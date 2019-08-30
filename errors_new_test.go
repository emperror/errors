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

func TestNewWithDetails(t *testing.T) {
	details := []interface{}{"key", "value"}
	err := NewWithDetails("error", details...)
	origErr := err.(*withDetails).error

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
			"%+v": {"error", "emperror.dev/errors.TestNewWithDetails\n\t.+/errors_new_test.go:52"},
		})
	})

	t.Run("details", func(t *testing.T) {
		d := err.(*withDetails).Details()

		for i, detail := range d {
			if got, want := detail, details[i]; got != want {
				t.Errorf("error detail does not match the expected one\nactual:   %+v\nexpected: %+v", got, want)
			}
		}
	})

	t.Run("details_missing_value", func(t *testing.T) {
		details := []interface{}{"key", nil}
		err := NewWithDetails("error", "key")

		d := err.(*withDetails).Details()

		for i, detail := range d {
			if got, want := detail, details[i]; got != want {
				t.Errorf("error detail does not match the expected one\nactual:   %+v\nexpected: %+v", got, want)
			}
		}
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
			"%+v": {"error: something went wrong", "emperror.dev/errors.TestErrorf\n\t.+/errors_new_test.go:99"},
		})
	})
}
