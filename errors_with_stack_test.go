package errors

import (
	"io"
	"testing"
)

func TestWithStack_Nil(t *testing.T) {
	err := WithStack(nil)

	if err != nil {
		t.Errorf("nil error is expected to result in nil\nactual:   %#v", err)
	}
}

func TestWithStack(t *testing.T) {
	origErr := NewPlain("msg")
	err := WithStack(origErr)

	testUnwrap(t, err, origErr)
	testFormatRegexp(t, 1, err, "%+v", "msg\nemperror.dev/errors.TestWithStack\n\t.+/errors_with_stack_test.go:18")
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
					"\t.+/errors_with_stack_test.go:41"},
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
					"\t.+/errors_with_stack_test.go:58"},
		},
		{
			WithStack(WithStack(io.EOF)),
			"%+v",
			[]string{"EOF",
				"emperror.dev/errors.TestWithStack_Format\n" +
					"\t.+/errors_with_stack_test.go:65",
				"emperror.dev/errors.TestWithStack_Format\n" +
					"\t.+/errors_with_stack_test.go:65"},
		},
	}

	for i, tt := range tests {
		testFormatCompleteCompare(t, i, tt.error, tt.format, tt.want, true)
	}
}

func TestWithStackDepth_Nil(t *testing.T) {
	err := WithStackDepth(nil, 0)

	if err != nil {
		t.Errorf("nil error is expected to result in nil\nactual:   %#v", err)
	}
}

func TestWithStackDepth(t *testing.T) {
	origErr := NewPlain("msg")
	err := WithStackDepth(origErr, 0)

	testUnwrap(t, err, origErr)
	testFormatRegexp(t, 1, err, "%+v", "msg\nemperror.dev/errors.TestWithStackDepth\n\t.+/errors_with_stack_test.go:90")
}

func TestWithStackDepth_CustomDepth(t *testing.T) {
	origErr := NewPlain("msg")
	var err error

	func() {
		err = WithStackDepth(origErr, 1)
	}()

	testUnwrap(t, err, origErr)
	testFormatRegexp(t, 1, err, "%+v", "msg\nemperror.dev/errors.TestWithStackDepth\n\t.+/errors_with_stack_test.go:102")
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
					"\t.+/errors_with_stack_test.go:125"},
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
					"\t.+/errors_with_stack_test.go:142"},
		},
		{
			WithStack(WithStack(io.EOF)),
			"%+v",
			[]string{"EOF",
				"emperror.dev/errors.TestWithStackDepth_Format\n" +
					"\t.+/errors_with_stack_test.go:149",
				"emperror.dev/errors.TestWithStackDepth_Format\n" +
					"\t.+/errors_with_stack_test.go:149"},
		},
	}

	for i, tt := range tests {
		testFormatCompleteCompare(t, i, tt.error, tt.format, tt.want, true)
	}
}
