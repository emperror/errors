package errors

import (
	"testing"
)

func TestWithMessage(t *testing.T) {
	origErr := NewPlain("something went wrong")
	err := WithMessage(origErr, "error")

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
			"%q":  {`error: something went wrong`}, // TODO: quotes?
			"%v":  {"error: something went wrong"},
			"%+v": {"something went wrong", "error"},
		})
	})

	t.Run("nil", func(t *testing.T) {
		checkErrorNil(t, WithMessage(nil, "error"))
	})
}

func TestWithMessagef(t *testing.T) {
	origErr := NewPlain("something went wrong")
	err := WithMessagef(origErr, "%s", "error")

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
			"%q":  {`error: something went wrong`}, // TODO: quotes?
			"%v":  {"error: something went wrong"},
			"%+v": {"something went wrong", "error"},
		})
	})

	t.Run("nil", func(t *testing.T) {
		checkErrorNil(t, WithMessagef(nil, "%s", "error"))
	})
}
