package errors_test

import (
	"fmt"
	"os"

	"emperror.dev/errors"
)

func ExampleAs() {
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *os.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}

	// Output:
	// Failed at path: non-existing
}

func ExampleIs() {
	var ErrSomething = errors.NewPlain("something")

	if err := errors.Wrap(ErrSomething, "error"); err != nil {
		if errors.Is(err, ErrSomething) {
			fmt.Println("something went wrong")
		} else {
			fmt.Println(err)
		}
	}

	// Output:
	// something went wrong
}

func ExampleUnwrap() {
	var ErrSomething = errors.NewPlain("something")

	if err := errors.WithMessage(ErrSomething, "error"); err != nil {
		if errors.Unwrap(err) == ErrSomething {
			fmt.Println("something went wrong")
		} else {
			fmt.Println(err)
		}
	}

	// Output:
	// something went wrong
}

func ExampleCause() {
	var ErrSomething = errors.NewPlain("something")

	if err := errors.Wrap(ErrSomething, "error"); err != nil {
		if errors.Cause(err) == ErrSomething {
			fmt.Println("something went wrong")
		} else {
			fmt.Println(err)
		}
	}

	// Output:
	// something went wrong
}
