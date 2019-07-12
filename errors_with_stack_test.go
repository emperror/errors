package errors

import (
	"testing"
)

func TestWithStack_Nil(t *testing.T) {
	err := WithStack(nil)

	if err != nil {
		t.Errorf("nil error is expected to result in nil\nactual:   %#v", err)
	}
}

func TestWithStack(t *testing.T) {
	origErr := NewPlain("something went wrong")
	err := WithStack(origErr)
	err2 := WithStack(err)

	t.Parallel()

	t.Run("error_message", func(t *testing.T) {
		checkErrorMessage(t, err, "something went wrong")
	})

	t.Run("unwrap", func(t *testing.T) {
		checkUnwrap(t, err, origErr)
	})

	t.Run("format", func(t *testing.T) {
		checkFormat(t, err, map[string][]string{
			"%s":  {"something went wrong"},
			"%q":  {`"something went wrong"`},
			"%v":  {"something went wrong"},
			"%+v": {"something went wrong", "emperror.dev/errors.TestWithStack\n\t.+/errors_with_stack_test.go:17"},
		})

		checkFormat(t, err2, map[string][]string{
			"%s": {"something went wrong"},
			"%q": {`"something went wrong"`},
			"%v": {"something went wrong"},
			"%+v": {
				"something went wrong",
				"emperror.dev/errors.TestWithStack\n\t.+/errors_with_stack_test.go:17",
				"emperror.dev/errors.TestWithStack\n\t.+/errors_with_stack_test.go:18",
			},
		})
	})

	t.Run("nil", func(t *testing.T) {
		checkErrorNil(t, WithStack(nil))
	})
}

func TestWithStackDepth(t *testing.T) {
	origErr := NewPlain("something went wrong")
	err := WithStackDepth(origErr, 0)
	err2 := WithStackDepth(err, 0)

	t.Parallel()

	t.Run("error_message", func(t *testing.T) {
		checkErrorMessage(t, err, "something went wrong")
	})

	t.Run("unwrap", func(t *testing.T) {
		checkUnwrap(t, err, origErr)
	})

	t.Run("format", func(t *testing.T) {
		checkFormat(t, err, map[string][]string{
			"%s":  {"something went wrong"},
			"%q":  {`"something went wrong"`},
			"%v":  {"something went wrong"},
			"%+v": {"something went wrong", "emperror.dev/errors.TestWithStackDepth\n\t.+/errors_with_stack_test.go:57"},
		})

		checkFormat(t, err2, map[string][]string{
			"%s": {"something went wrong"},
			"%q": {`"something went wrong"`},
			"%v": {"something went wrong"},
			"%+v": {
				"something went wrong",
				"emperror.dev/errors.TestWithStackDepth\n\t.+/errors_with_stack_test.go:57",
				"emperror.dev/errors.TestWithStackDepth\n\t.+/errors_with_stack_test.go:58",
			},
		})
	})

	t.Run("nil", func(t *testing.T) {
		checkErrorNil(t, WithStack(nil))
	})
}

func TestWithStackDepth_CustomDepth(t *testing.T) {
	origErr := NewPlain("something went wrong")

	var err, err2 error

	func() {
		err = WithStackDepth(origErr, 1)
		err2 = WithStackDepth(err, 1)
	}()

	t.Parallel()

	t.Run("error_message", func(t *testing.T) {
		checkErrorMessage(t, err, "something went wrong")
	})

	t.Run("unwrap", func(t *testing.T) {
		checkUnwrap(t, err, origErr)
	})

	t.Run("format", func(t *testing.T) {
		checkFormat(t, err, map[string][]string{
			"%s": {"something went wrong"},
			"%q": {`"something went wrong"`},
			"%v": {"something went wrong"},
			"%+v": {
				"something went wrong",
				"emperror.dev/errors.TestWithStackDepth_CustomDepth\n\t.+/errors_with_stack_test.go:103",
			},
		})

		checkFormat(t, err2, map[string][]string{
			"%s": {"something went wrong"},
			"%q": {`"something went wrong"`},
			"%v": {"something went wrong"},
			"%+v": {
				"something went wrong",
				"emperror.dev/errors.TestWithStackDepth_CustomDepth\n\t.+/errors_with_stack_test.go:103",
				"emperror.dev/errors.TestWithStackDepth_CustomDepth\n\t.+/errors_with_stack_test.go:103",
			},
		})
	})

	t.Run("nil", func(t *testing.T) {
		checkErrorNil(t, WithStack(nil))
	})
}
