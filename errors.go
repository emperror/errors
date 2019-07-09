// Package errors is a drop-in replacement for the standard errors package and github.com/pkg/errors.
package errors

import (
	"fmt"
	"io"
)

// New returns a new error annotated with stack trace.
func New(message string) error {
	return WithStackDepth(NewPlain(message), 1)
}

// Errorf returns a new error with a formatted message and annotated with stack trace.
func Errorf(format string, a ...interface{}) error {
	return WithStackDepth(NewPlain(fmt.Sprintf(format, a...)), 1)
}

// NewPlain returns a simple error without any annotated context, like stack trace.
// Useful for creating sentinel errors and in testing.
func NewPlain(message string) error {
	return &plainError{message}
}

// plainError is a trivial implementation of error.
type plainError struct {
	msg string
}

func (e *plainError) Error() string {
	return e.msg
}

// WithStack annotates err with a stack trace at the point WithStack was called.
// If err is nil, WithStack returns nil.
func WithStack(err error) error {
	return WithStackDepth(err, 1)
}

// WithStackDepth annotates err with a stack trace at the given call depth.
// Zero identifies the caller of WithStackDepth itself.
// If err is nil, WithStackDepth returns nil.
func WithStackDepth(err error, depth int) error {
	if err == nil {
		return nil
	}

	return &withStack{
		err,
		callers(depth + 1),
	}
}

type withStack struct {
	error
	*stack
}

func (w *withStack) Cause() error  { return w.error }
func (w *withStack) Unwrap() error { return w.error }

// nolint: errcheck
func (w *withStack) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", w.Cause())
			w.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, w.Error())
	case 'q':
		fmt.Fprintf(s, "%q", w.Error())
	}
}
