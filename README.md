# Emperror: Errors

[![CircleCI](https://circleci.com/gh/emperror/errors.svg?style=svg)](https://circleci.com/gh/emperror/errors)
[![Go Report Card](https://goreportcard.com/badge/emperror.dev/errors?style=flat-square)](https://goreportcard.com/report/emperror.dev/errors)
[![GolangCI](https://golangci.com/badges/github.com/emperror/errors.svg)](https://golangci.com/r/github.com/emperror/errors)
[![Go Version](https://img.shields.io/badge/go%20version-%3E=1.12-orange.svg?style=flat-square)](https://github.com/emperror/errors)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/emperror.dev/errors)

**Drop-in replacement for the standard library `errors` package and [github.com/pkg/errors](https://github.com/pkg/errors).**

This is a single, lightweight library merging the features of standard library `errors` package
and [github.com/pkg/errors](https://github.com/pkg/errors). It also backports a few features
(like Go 1.13 error handling related features).

Standard library features:
- `New` creates an error with stack trace
- `Unwrap` supports both Go 1.13 wrapper (`interface { Unwrap() error }`) and **pkg/errors** causer (`interface { Cause() error }`) interface
- Backported `Is` and `As` functions

[github.com/pkg/errors](https://github.com/pkg/errors) features:
- `New`, `Errorf`, `WithMessage`, `WithMessagef`, `WithStack`, `Wrap`, `Wrapf` functions behave the same way as in the original library
- `Cause` supports both Go 1.13 wrapper (`interface { Unwrap() error }`) and **pkg/errors** causer (`interface { Cause() error }`) interface

Additional features:
- `NewPlain` creates a new error without any attached context, like stack trace
- `WithStackDepth` allows attaching stack trace with a custom caller depth
- `WithStackDepthIf`, `WithStackIf`, `WrapIf`, `WrapIff` only annotate errors with a stack trace if there isn't one already in the error chain
- Multi error aggregating multiple errors into a single value
- `NewWithDetails`, `WithDetails` and `Wrap*WithDetails` functions to add key-value pairs to an error


## Installation

```bash
go get emperror.dev/errors
```


## Usage

```go
package main

import "emperror.dev/errors"

// ErrSomethingWentWrong is a sentinel error which can be useful within a single API layer.
var ErrSomethingWentWrong = errors.NewPlain("something went wrong")

// ErrMyError is an error that can be returned from a public API.
type ErrMyError struct {
	Msg string
}

func (e ErrMyError) Error() string {
	return e.Msg
}

func foo() error {
	// Attach stack trace to the sentinel error.
	return errors.WithStack(ErrSomethingWentWrong)
}

func bar() error {
	return errors.Wrap(ErrMyError{"something went wrong"}, "error")
}

func main() {
	if err := foo(); err != nil {
		if errors.Cause(err) == ErrSomethingWentWrong { // or errors.Is(ErrSomethingWentWrong)
			// handle error
		}
	}
	
	if err := bar(); err != nil {
		if errors.As(err, &ErrMyError{}) {
			// handle error
		}
	}
}
```


## Development

When all coding and testing is done, please run the test suite:

``` bash
$ make check
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.

Certain parts of this library are inspired by (or entirely copied from) various third party libraries.
Their licenses can be found in the [Third Party License File](LICENSE_THIRD_PARTY).
