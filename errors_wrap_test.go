package errors

import (
	"testing"
)

func TestWrap_Nil(t *testing.T) {
	err := Wrap(nil, "error")

	if err != nil {
		t.Errorf("nil error is expected to result in nil\nactual:   %#v", err)
	}
}

func TestWrap(t *testing.T) {
	origErr := NewPlain("something went wrong")
	err := Wrap(origErr, "error")

	t.Parallel()

	t.Run("error_message", func(t *testing.T) {
		checkErrorMessage(t, err, "error: something went wrong")
	})

	t.Run("unwrap", func(t *testing.T) {
		checkUnwrap(t, err, err.(*withStack).error)
	})

	t.Run("format", func(t *testing.T) {
		checkFormat(t, err, map[string][]string{
			"%s":  {"error: something went wrong"},
			"%q":  {`"error: something went wrong"`},
			"%v":  {"error: something went wrong"},
			"%+v": {"something went wrong", "error", "emperror.dev/errors.TestWrap\n\t.+/errors_wrap_test.go:17"},
		})
	})

	t.Run("nil", func(t *testing.T) {
		checkErrorNil(t, Wrap(nil, "error"))
	})
}

func TestWrapf(t *testing.T) {
	origErr := NewPlain("something went wrong")
	err := Wrapf(origErr, "%s", "error")
	err2 := Wrapf(err, "%s", "panic")

	t.Parallel()

	t.Run("error_message", func(t *testing.T) {
		checkErrorMessage(t, err, "error: something went wrong")
	})

	t.Run("unwrap", func(t *testing.T) {
		checkUnwrap(t, err, err.(*withStack).error)
	})

	t.Run("format", func(t *testing.T) {
		checkFormat(t, err, map[string][]string{
			"%s":  {"error: something went wrong"},
			"%q":  {`"error: something went wrong"`},
			"%v":  {"error: something went wrong"},
			"%+v": {"something went wrong", "error", "emperror.dev/errors.TestWrapf\n\t.+/errors_wrap_test.go:45"},
		})

		checkFormat(t, err2, map[string][]string{
			"%s": {"panic: error: something went wrong"},
			"%q": {`"panic: error: something went wrong"`},
			"%v": {"panic: error: something went wrong"},
			"%+v": {
				"something went wrong",
				"error",
				"emperror.dev/errors.TestWrapf\n\t.+/errors_wrap_test.go:45",
				"panic",
				"emperror.dev/errors.TestWrapf\n\t.+/errors_wrap_test.go:46",
			},
		})
	})

	t.Run("nil", func(t *testing.T) {
		checkErrorNil(t, Wrapf(nil, "%s", "error"))
	})
}
