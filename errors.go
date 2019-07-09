// Package errors is a drop-in replacement for the standard errors package and github.com/pkg/errors.
package errors

import (
	"fmt"
	"io"
)

// New returns a new error annotated with stack trace at the point New is called.
func New(message string) error {
	return WithStackDepth(NewPlain(message), 1)
}

// Errorf returns a new error with a formatted message and annotated with stack trace at the point Errorf is called.
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
		error: err,
		stack: callers(depth + 1),
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
			fmt.Fprintf(s, "%+v", w.error)
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

// WithMessage annotates err with a new message.
// If err is nil, WithMessage returns nil.
func WithMessage(err error, message string) error {
	if err == nil {
		return nil
	}

	return &withMessage{
		error: err,
		msg:   message,
	}
}

// WithMessagef annotates err with the format specifier.
// If err is nil, WithMessagef returns nil.
func WithMessagef(err error, format string, a ...interface{}) error {
	if err == nil {
		return nil
	}

	return &withMessage{
		error: err,
		msg:   fmt.Sprintf(format, a...),
	}
}

type withMessage struct {
	error error
	msg   string
}

func (w *withMessage) Error() string { return w.msg + ": " + w.error.Error() }
func (w *withMessage) Cause() error  { return w.error }
func (w *withMessage) Unwrap() error { return w.error }

// nolint: errcheck
func (w *withMessage) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v\n", w.error)
			io.WriteString(s, w.msg)
			return
		}
		fallthrough
	case 's', 'q':
		io.WriteString(s, w.Error())
	}
}

// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called, and the supplied message.
// If err is nil, Wrap returns nil.
func Wrap(err error, message string) error {
	return WithStackDepth(WithMessage(err, message), 1)
}

// Wrapf returns an error annotating err with a stack trace
// at the point Wrapf is called, and the format specifier.
// If err is nil, Wrapf returns nil.
func Wrapf(err error, format string, a ...interface{}) error {
	return WithStackDepth(WithMessagef(err, format, a...), 1)
}
