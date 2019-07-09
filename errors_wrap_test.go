package errors

import (
	"io"
	"testing"
)

func TestWrap_Nil(t *testing.T) {
	err := Wrap(nil, "error")

	if err != nil {
		t.Errorf("nil error is expected to result in nil\nactual:   %#v", err)
	}
}

func TestWrap(t *testing.T) {
	origErr := NewPlain("msg")
	err := Wrap(origErr, "error")

	if got, want := err.Error(), "error: msg"; got != want {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}

	// TODO: test root cause?
	// testUnwrap(t, err, origErr)
	testFormatRegexp(t, 1, err, "%+v", "msg\nerror\nemperror.dev/errors.TestWrap\n\t.+/errors_wrap_test.go:18")
}

func TestWrap_Format(t *testing.T) {
	tests := []struct {
		error
		format string
		want   []string
	}{
		{
			Wrap(io.EOF, "error"),
			"%s",
			[]string{"error: EOF"},
		},
		{
			Wrap(io.EOF, "error"),
			"%v",
			[]string{"error: EOF"},
		},
		{
			Wrap(io.EOF, "error"),
			"%+v",
			[]string{"EOF", "error",
				"emperror.dev/errors.TestWrap_Format\n" +
					"\t.+/errors_wrap_test.go:46"},
		},
		{
			Wrap(NewPlain("error"), "error 2"),
			"%s",
			[]string{"error 2: error"},
		},
		{
			Wrap(NewPlain("error"), "error 2"),
			"%v",
			[]string{"error 2: error"},
		},
		{
			Wrap(NewPlain("error"), "error 2"),
			"%+v",
			[]string{"error", "error 2",
				"emperror.dev/errors.TestWrap_Format\n" +
					"\t.+/errors_wrap_test.go:63"},
		},
		{
			Wrap(Wrap(io.EOF, "error 2"), "error 3"),
			"%+v",
			[]string{"EOF", "error 2",
				"emperror.dev/errors.TestWrap_Format\n" +
					"\t.+/errors_wrap_test.go:70",
					"error 3",
				"emperror.dev/errors.TestWrap_Format\n" +
					"\t.+/errors_wrap_test.go:70"},
		},
	}

	for i, tt := range tests {
		testFormatCompleteCompare(t, i, tt.error, tt.format, tt.want, true)
	}
}

func TestWrapf_Nil(t *testing.T) {
	err := Wrapf(nil, "error %d", 1)

	if err != nil {
		t.Errorf("nil error is expected to result in nil\nactual:   %#v", err)
	}
}

func TestWrapf(t *testing.T) {
	origErr := NewPlain("msg")
	err := Wrapf(origErr, "error %d", 1)

	if got, want := err.Error(), "error 1: msg"; got != want {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}

	// TODO: test root cause?
	// testUnwrap(t, err, origErr)
	testFormatRegexp(t, 1, err, "%+v", "msg\nerror\nemperror.dev/errors.TestWrapf\n\t.+/errors_wrap_test.go:96")
}

func TestWrapf_Format(t *testing.T) {
	tests := []struct {
		error
		format string
		want   []string
	}{
		{
			Wrapf(io.EOF, "error %d", 1),
			"%s",
			[]string{"error 1: EOF"},
		},
		{
			Wrapf(io.EOF, "error %d", 1),
			"%v",
			[]string{"error 1: EOF"},
		},
		{
			Wrapf(io.EOF, "error %d", 1),
			"%+v",
			[]string{"EOF", "error 1",
				"emperror.dev/errors.TestWrapf_Format\n" +
					"\t.+/errors_wrap_test.go:124"},
		},
		{
			Wrapf(NewPlain("error"), "error %d", 2),
			"%s",
			[]string{"error 2: error"},
		},
		{
			Wrapf(NewPlain("error"), "error %d", 2),
			"%v",
			[]string{"error 2: error"},
		},
		{
			Wrapf(NewPlain("error"), "error %d", 2),
			"%+v",
			[]string{"error", "error 2",
				"emperror.dev/errors.TestWrapf_Format\n" +
					"\t.+/errors_wrap_test.go:141"},
		},
		{
			Wrapf(Wrapf(io.EOF, "error %d", 2), "error %d", 3),
			"%+v",
			[]string{"EOF", "error 2",
				"emperror.dev/errors.TestWrapf_Format\n" +
					"\t.+/errors_wrap_test.go:148",
				"error 3",
				"emperror.dev/errors.TestWrapf_Format\n" +
					"\t.+/errors_wrap_test.go:148"},
		},
	}

	for i, tt := range tests {
		testFormatCompleteCompare(t, i, tt.error, tt.format, tt.want, true)
	}
}
