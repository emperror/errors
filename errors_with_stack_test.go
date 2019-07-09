package errors

import (
	"io"
	"testing"
)

func TestWithStackNil(t *testing.T) {
	err := WithStack(nil)
	if err != nil {
		t.Errorf("nil error is expected to result in nil\nactual:   %#v", err)
	}
}

func TestWithStack(t *testing.T) {
	origErr := NewPlain("msg")
	err := WithStack(origErr)

	testUnwrap(t, err, origErr)
	testFormatRegexp(t, 1, err, "%+v", "msg\nemperror.dev/errors.TestWithStack\n\t.+/errors_with_stack_test.go:17")
}

func TestWithStack_Format(t *testing.T) {
	tests := []struct {
		error
		format string
		want   []string
	}{
		{
			WithStack(io.EOF),
			"%s",
			[]string{"EOF"},
		},
		{
			WithStack(io.EOF),
			"%v",
			[]string{"EOF"},
		},
		{
			WithStack(io.EOF),
			"%+v",
			[]string{"EOF",
				"emperror.dev/errors.TestWithStack_Format\n" +
					"\t.+/errors_with_stack_test.go:40"},
		},
		{
			WithStack(NewPlain("error")),
			"%s",
			[]string{"error"},
		},
		{
			WithStack(NewPlain("error")),
			"%v",
			[]string{"error"},
		},
		{
			WithStack(NewPlain("error")),
			"%+v",
			[]string{"error",
				"emperror.dev/errors.TestWithStack_Format\n" +
					"\t.+/errors_with_stack_test.go:57"},
		},
		{
			WithStack(WithStack(io.EOF)),
			"%+v",
			[]string{"EOF",
				"emperror.dev/errors.TestWithStack_Format\n" +
					"\t.+/errors_with_stack_test.go:64",
				"emperror.dev/errors.TestWithStack_Format\n" +
					"\t.+/errors_with_stack_test.go:64"},
		},
	}

	for i, tt := range tests {
		testFormatCompleteCompare(t, i, tt.error, tt.format, tt.want, true)
	}
}

func TestWithStackDepthNil(t *testing.T) {
	err := WithStackDepth(nil, 0)
	if err != nil {
		t.Errorf("nil error is expected to result in nil\nactual:   %#v", err)
	}
}

func TestWithStackDepth(t *testing.T) {
	origErr := NewPlain("msg")
	err := WithStackDepth(origErr, 0)

	testUnwrap(t, err, origErr)
	testFormatRegexp(t, 1, err, "%+v", "msg\nemperror.dev/errors.TestWithStackDepth\n\t.+/errors_with_stack_test.go:88")
}

func TestWithStackDepth_CustomDepth(t *testing.T) {
	origErr := NewPlain("msg")
	var err error

	func() {
		err = WithStackDepth(origErr, 1)
	}()

	testUnwrap(t, err, origErr)
	testFormatRegexp(t, 1, err, "%+v", "msg\nemperror.dev/errors.TestWithStackDepth\n\t.+/errors_with_stack_test.go:100")
}

func TestWithStackDepth_Format(t *testing.T) {
	tests := []struct {
		error
		format string
		want   []string
	}{
		{
			WithStack(io.EOF),
			"%s",
			[]string{"EOF"},
		},
		{
			WithStack(io.EOF),
			"%v",
			[]string{"EOF"},
		},
		{
			WithStack(io.EOF),
			"%+v",
			[]string{"EOF",
				"emperror.dev/errors.TestWithStackDepth_Format\n" +
					"\t.+/errors_with_stack_test.go:123"},
		},
		{
			WithStack(NewPlain("error")),
			"%s",
			[]string{"error"},
		},
		{
			WithStack(NewPlain("error")),
			"%v",
			[]string{"error"},
		},
		{
			WithStack(NewPlain("error")),
			"%+v",
			[]string{"error",
				"emperror.dev/errors.TestWithStackDepth_Format\n" +
					"\t.+/errors_with_stack_test.go:140"},
		},
		{
			WithStack(WithStack(io.EOF)),
			"%+v",
			[]string{"EOF",
				"emperror.dev/errors.TestWithStackDepth_Format\n" +
					"\t.+/errors_with_stack_test.go:147",
				"emperror.dev/errors.TestWithStackDepth_Format\n" +
					"\t.+/errors_with_stack_test.go:147"},
		},
	}

	for i, tt := range tests {
		testFormatCompleteCompare(t, i, tt.error, tt.format, tt.want, true)
	}
}
