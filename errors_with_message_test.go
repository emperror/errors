package errors

import (
	"io"
	"testing"
)

func TestWithMessage_Nil(t *testing.T) {
	err := WithMessage(nil, "error")

	if err != nil {
		t.Errorf("nil error is expected to result in nil\nactual:   %#v", err)
	}
}

func TestWithMessage(t *testing.T) {
	origErr := NewPlain("msg")
	err := WithMessage(origErr, "error")

	if got, want := err.Error(), "error: msg"; got != want {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

func TestWithMessage_Format(t *testing.T) {
	tests := []struct {
		error
		format string
		want   []string
	}{
		{
			WithMessage(io.EOF, "error"),
			"%s",
			[]string{"error: EOF"},
		},
		{
			WithMessage(io.EOF, "error"),
			"%v",
			[]string{"error: EOF"},
		},
		{
			WithMessage(io.EOF, "error"),
			"%+v",
			[]string{"EOF", "error"},
		},
		{
			WithMessage(NewPlain("error"), "error 2"),
			"%s",
			[]string{"error 2: error"},
		},
		{
			WithMessage(NewPlain("error"), "error 2"),
			"%v",
			[]string{"error 2: error"},
		},
		{
			WithMessage(NewPlain("error"), "error 2"),
			"%+v",
			[]string{"error", "error 2"},
		},
	}

	for i, tt := range tests {
		testFormatCompleteCompare(t, i, tt.error, tt.format, tt.want, true)
	}
}

func TestWithMessagef_Nil(t *testing.T) {
	err := WithMessagef(nil, "error %d", 1)

	if err != nil {
		t.Errorf("nil error is expected to result in nil\nactual:   %#v", err)
	}
}

func TestWithMessagef(t *testing.T) {
	origErr := NewPlain("msg")
	err := WithMessagef(origErr, "error %d", 1)

	if got, want := err.Error(), "error 1: msg"; got != want {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

func TestWithMessagef_Format(t *testing.T) {
	tests := []struct {
		error
		format string
		want   []string
	}{
		{
			WithMessagef(io.EOF, "error %d", 1),
			"%s",
			[]string{"error 1: EOF"},
		},
		{
			WithMessagef(io.EOF, "error %d", 1),
			"%v",
			[]string{"error 1: EOF"},
		},
		{
			WithMessagef(io.EOF, "error %d", 1),
			"%+v",
			[]string{"EOF", "error 1"},
		},
		{
			WithMessagef(NewPlain("error"), "error %d", 2),
			"%s",
			[]string{"error 2: error"},
		},
		{
			WithMessagef(NewPlain("error"), "error %d", 2),
			"%v",
			[]string{"error 2: error"},
		},
		{
			WithMessagef(NewPlain("error"), "error %d", 2),
			"%+v",
			[]string{"error", "error 2"},
		},
	}

	for i, tt := range tests {
		testFormatCompleteCompare(t, i, tt.error, tt.format, tt.want, true)
	}
}
