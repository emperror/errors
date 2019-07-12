package errors

import (
	"testing"
)

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
			"%+v": {"something went wrong", "emperror.dev/errors.TestWithStack\n\t.+/errors_with_stack_test.go:9"},
		})

		checkFormat(t, err2, map[string][]string{
			"%s": {"something went wrong"},
			"%q": {`"something went wrong"`},
			"%v": {"something went wrong"},
			"%+v": {
				"something went wrong",
				"emperror.dev/errors.TestWithStack\n\t.+/errors_with_stack_test.go:9",
				"emperror.dev/errors.TestWithStack\n\t.+/errors_with_stack_test.go:10",
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
			"%+v": {"something went wrong", "emperror.dev/errors.TestWithStackDepth\n\t.+/errors_with_stack_test.go:49"},
		})

		checkFormat(t, err2, map[string][]string{
			"%s": {"something went wrong"},
			"%q": {`"something went wrong"`},
			"%v": {"something went wrong"},
			"%+v": {
				"something went wrong",
				"emperror.dev/errors.TestWithStackDepth\n\t.+/errors_with_stack_test.go:49",
				"emperror.dev/errors.TestWithStackDepth\n\t.+/errors_with_stack_test.go:50",
			},
		})
	})

	t.Run("nil", func(t *testing.T) {
		checkErrorNil(t, WithStackDepth(nil, 0))
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
				"emperror.dev/errors.TestWithStackDepth_CustomDepth\n\t.+/errors_with_stack_test.go:95",
			},
		})

		checkFormat(t, err2, map[string][]string{
			"%s": {"something went wrong"},
			"%q": {`"something went wrong"`},
			"%v": {"something went wrong"},
			"%+v": {
				"something went wrong",
				"emperror.dev/errors.TestWithStackDepth_CustomDepth\n\t.+/errors_with_stack_test.go:95",
				"emperror.dev/errors.TestWithStackDepth_CustomDepth\n\t.+/errors_with_stack_test.go:95",
			},
		})
	})

	t.Run("nil", func(t *testing.T) {
		checkErrorNil(t, WithStackDepth(nil, 0))
	})
}

func TestWithStackIf(t *testing.T) {
	origErr := NewPlain("something went wrong")
	err := WithStackIf(origErr)
	err2 := WithStackIf(err)

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
			"%+v": {"something went wrong", "emperror.dev/errors.TestWithStackIf\n\t.+/errors_with_stack_test.go:137"},
		})

		checkFormat(t, err2, map[string][]string{
			"%s":  {"something went wrong"},
			"%q":  {`"something went wrong"`},
			"%v":  {"something went wrong"},
			"%+v": {"something went wrong", "emperror.dev/errors.TestWithStackIf\n\t.+/errors_with_stack_test.go:137"},
		})
	})

	t.Run("nil", func(t *testing.T) {
		checkErrorNil(t, WithStackIf(nil))
	})
}

func TestWithStackDepthIf(t *testing.T) {
	origErr := NewPlain("something went wrong")
	err := WithStackDepthIf(origErr, 0)
	err2 := WithStackDepthIf(err, 0)

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
			"%+v": {"something went wrong", "emperror.dev/errors.TestWithStackDepthIf\n\t.+/errors_with_stack_test.go:173"},
		})

		checkFormat(t, err2, map[string][]string{
			"%s":  {"something went wrong"},
			"%q":  {`"something went wrong"`},
			"%v":  {"something went wrong"},
			"%+v": {"something went wrong", "emperror.dev/errors.TestWithStackDepthIf\n\t.+/errors_with_stack_test.go:173"},
		})
	})

	t.Run("nil", func(t *testing.T) {
		checkErrorNil(t, WithStack(nil))
	})
}
