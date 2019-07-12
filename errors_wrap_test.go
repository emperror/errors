package errors

import (
	"testing"
)

func TestWrap(t *testing.T) {
	origErr := NewPlain("something went wrong")
	err := Wrap(origErr, "error")
	err2 := Wrap(err, "panic")

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
			"%+v": {"something went wrong", "error", "emperror.dev/errors.TestWrap\n\t.+/errors_wrap_test.go:9"},
		})

		checkFormat(t, err2, map[string][]string{
			"%s": {"panic: error: something went wrong"},
			"%q": {`"panic: error: something went wrong"`},
			"%v": {"panic: error: something went wrong"},
			"%+v": {
				"something went wrong",
				"error",
				"emperror.dev/errors.TestWrap\n\t.+/errors_wrap_test.go:9",
				"panic",
				"emperror.dev/errors.TestWrap\n\t.+/errors_wrap_test.go:10",
			},
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
			"%+v": {"something went wrong", "error", "emperror.dev/errors.TestWrapf\n\t.+/errors_wrap_test.go:51"},
		})

		checkFormat(t, err2, map[string][]string{
			"%s": {"panic: error: something went wrong"},
			"%q": {`"panic: error: something went wrong"`},
			"%v": {"panic: error: something went wrong"},
			"%+v": {
				"something went wrong",
				"error",
				"emperror.dev/errors.TestWrapf\n\t.+/errors_wrap_test.go:51",
				"panic",
				"emperror.dev/errors.TestWrapf\n\t.+/errors_wrap_test.go:52",
			},
		})
	})

	t.Run("nil", func(t *testing.T) {
		checkErrorNil(t, Wrapf(nil, "%s", "error"))
	})
}

func TestWrapIf(t *testing.T) {
	origErr := NewPlain("something went wrong")
	err := WrapIf(origErr, "error")
	err2 := WrapIf(err, "panic")

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
			"%+v": {"something went wrong", "error", "emperror.dev/errors.TestWrapIf\n\t.+/errors_wrap_test.go:93"},
		})

		checkFormat(t, err2, map[string][]string{
			"%s": {"panic: error: something went wrong"},
			"%q": {`panic: error: something went wrong`}, // TODO: quotes?
			"%v": {"panic: error: something went wrong"},
			"%+v": {
				"something went wrong",
				"error",
				"emperror.dev/errors.TestWrapIf\n\t.+/errors_wrap_test.go:93",
				"panic",
			},
		})
	})

	t.Run("nil", func(t *testing.T) {
		checkErrorNil(t, WrapIf(nil, "error"))
	})
}

func TestWrapIff(t *testing.T) {
	origErr := NewPlain("something went wrong")
	err := WrapIff(origErr, "%s", "error")
	err2 := WrapIff(err, "%s", "panic")

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
			"%+v": {"something went wrong", "error", "emperror.dev/errors.TestWrapIff\n\t.+/errors_wrap_test.go:134"},
		})

		checkFormat(t, err2, map[string][]string{
			"%s": {"panic: error: something went wrong"},
			"%q": {`panic: error: something went wrong`}, // TODO: quotes?
			"%v": {"panic: error: something went wrong"},
			"%+v": {
				"something went wrong",
				"error",
				"emperror.dev/errors.TestWrapIff\n\t.+/errors_wrap_test.go:134",
				"panic",
			},
		})
	})

	t.Run("nil", func(t *testing.T) {
		checkErrorNil(t, WrapIff(nil, "%s", "error"))
	})
}
