package errors_test

import (
	"fmt"
	"os"

	"emperror.dev/errors"
)

// nolint: unused
func ExampleSentinel() {
	const ErrSomething = errors.Sentinel("something went wrong")

	// Output:
}

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

func ExampleCombine() {
	err := errors.Combine(
		errors.NewPlain("call 1 failed"),
		nil, // successful request
		errors.NewPlain("call 3 failed"),
		nil, // successful request
		errors.NewPlain("call 5 failed"),
	)
	fmt.Printf("%+v", err)
	// Output:
	// the following errors occurred:
	//  -  call 1 failed
	//  -  call 3 failed
	//  -  call 5 failed
}

func ExampleCombine_loop() {
	var errs []error

	for i := 1; i < 6; i++ {
		if i%2 == 0 {
			continue
		}

		err := errors.NewPlain(fmt.Sprintf("call %d failed", i))
		errs = append(errs, err)
	}
	err := errors.Combine(errs...)
	fmt.Printf("%+v", err)
	// Output:
	// the following errors occurred:
	//  -  call 1 failed
	//  -  call 3 failed
	//  -  call 5 failed
}

func ExampleAppend() {
	var err error
	err = errors.Append(err, errors.NewPlain("call 1 failed"))
	err = errors.Append(err, errors.NewPlain("call 2 failed"))
	fmt.Println(err)
	// Output:
	// call 1 failed; call 2 failed
}

func ExampleGetErrors() {
	err := errors.Combine(
		nil, // successful request
		errors.NewPlain("call 2 failed"),
		errors.NewPlain("call 3 failed"),
	)
	err = errors.Append(err, nil) // successful request
	err = errors.Append(err, errors.NewPlain("call 5 failed"))

	errs := errors.GetErrors(err)
	for _, err := range errs {
		fmt.Println(err)
	}
	// Output:
	// call 2 failed
	// call 3 failed
	// call 5 failed
}
