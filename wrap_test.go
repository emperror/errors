// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code (or at least parts of it) is governed by a BSD-style
// license that can be found in the LICENSE_THIRD_PARTY file.

package errors_test

import (
	"io"
	"reflect"
	"testing"

	"emperror.dev/errors"
)

func TestUnwrap(t *testing.T) {
	err1 := errors.NewPlain("1")
	erra := wrapped{"wrap 2", err1}

	testCases := []struct {
		err  error
		want error
	}{
		{nil, nil},
		{wrapped{"wrapped", nil}, nil},
		{caused{"wrapped", nil}, nil},
		{err1, nil},
		{erra, err1},
		{wrapped{"wrap 3", erra}, erra},
		{caused{"wrap 3", erra}, erra},
	}
	for _, tc := range testCases {
		if got := errors.Unwrap(tc.err); got != tc.want {
			t.Errorf("Unwrap(%v) = %v, want %v", tc.err, got, tc.want)
		}
	}
}

type wrapped struct {
	msg string
	err error
}

func (e wrapped) Error() string { return e.msg }

func (e wrapped) Unwrap() error { return e.err }

type caused struct {
	msg string
	err error
}

func (e caused) Error() string { return e.msg }

func (e caused) Cause() error { return e.err }

func TestUnwrapEach(t *testing.T) {
	err := errors.WithMessage(
		errors.WithMessage(
			errors.WithMessage(
				errors.New("level 0"),
				"level 1",
			),
			"level 2",
		),
		"level 3",
	)

	var i int
	fn := func(err error) bool {
		i++

		return true
	}

	errors.UnwrapEach(err, fn)

	if got, want := i, 5; got != want {
		t.Errorf("error chain length does not match the expected one\nactual:   %d\nexpected: %d", got, want)
	}
}

func TestUnwrapEach_BreakTheLoop(t *testing.T) {
	err := errors.WithMessage(
		errors.WithMessage(
			errors.WithMessage(
				errors.New("level 0"),
				"level 1",
			),
			"level 2",
		),
		"level 3",
	)

	var i int
	fn := func(err error) bool {
		i++

		return !(i > 2)
	}

	errors.UnwrapEach(err, fn)

	if got, want := i, 3; got != want {
		t.Errorf("error chain length does not match the expected one\nactual:   %d\nexpected: %d", got, want)
	}
}

func TestUnwrapEach_NilError(t *testing.T) {
	var i int
	fn := func(err error) bool {
		i++

		return !(i > 2)
	}

	errors.UnwrapEach(nil, fn)

	if got, want := i, 0; got != want {
		t.Errorf("error chain length does not match the expected one\nactual:   %d\nexpected: %d", got, want)
	}
}

type nilError struct{}

func (nilError) Error() string { return "nil error" }

func TestCause(t *testing.T) {
	x := errors.NewPlain("error")
	tests := []struct {
		err  error
		want error
	}{
		{
			// nil error is nil
			err:  nil,
			want: nil,
		},
		{
			// explicit nil error is nil
			err:  (error)(nil),
			want: nil,
		},
		{
			// typed nil is nil
			err:  (*nilError)(nil),
			want: (*nilError)(nil),
		},
		{
			// uncaused error is unaffected
			err:  io.EOF,
			want: io.EOF,
		},
		{
			// wrapped error returns cause
			err:  wrapped{"ignored", io.EOF},
			want: io.EOF,
		},
		{
			// caused error returns cause
			err:  caused{"ignored", io.EOF},
			want: io.EOF,
		},
		{
			err:  x, // return from errors.New
			want: x,
		},
		{
			errors.WithMessage(nil, "whoops"),
			nil,
		},
		{
			errors.WithMessage(io.EOF, "whoops"),
			io.EOF,
		},
		{
			errors.WithMessagef(nil, "%s", "whoops"),
			nil,
		},
		{
			errors.WithMessagef(io.EOF, "%s", "whoops"),
			io.EOF,
		},
		{
			errors.WithStack(nil),
			nil,
		},
		{
			errors.WithStack(io.EOF),
			io.EOF,
		},
		{
			errors.WithStackDepth(nil, 0),
			nil,
		},
		{
			errors.WithStackDepth(io.EOF, 0),
			io.EOF,
		},
	}

	for i, tt := range tests {
		got := errors.Cause(tt.err)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("test %d: got %#v, want %#v", i+1, got, tt.want)
		}
	}
}
