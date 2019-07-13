package errors

// multiError aggregates multiple errors into a single value.
// Also implements the error interface so it can be returned as an error.
type multiError struct {
	errors []error
	msg    string
}

// Error implements the error interface.
func (e *multiError) Error() string {
	if e == nil {
		return ""
	}

	if e.msg != "" {
		return e.msg
	}

	return "multiple errors happened"
}

// Errors returns the list of wrapped errors.
func (e *multiError) Errors() []error {
	if e == nil {
		return nil
	}

	return e.errors
}

// createMultiError aggregates the given list of errors into a single error in an efficient way.
// Optimizations include:
// 	- flattening multiErrors
// 	- calculating the necessary size for the error slice
// 	- skipping nil errors
func createMultiError(errors []error) error {
	first := true
	var count, firstIdx int
	isFlat := true

	for i, err := range errors {
		if err == nil {
			continue
		}

		count++
		if first {
			first = false
			firstIdx = i
		}

		if m, ok := err.(*multiError); ok {
			count += len(m.errors) - 1 // the error itself is already counted once
			isFlat = false
		}
	}

	// no errors, return nil
	if count == 0 {
		return nil
	}

	// there is one error, identified by firstIdx, return it
	if count == 1 && isFlat {
		return errors[firstIdx]
	}

	if count == len(errors) && isFlat {
		// TODO: copy the errors?
		return &multiError{errors: errors}
	}

	errs := make([]error, 0, count)
	for _, err := range errors[firstIdx:] {
		if err == nil {
			continue
		}

		if m, ok := err.(*multiError); ok {
			errs = append(errs, m.errors...)
		} else {
			errs = append(errs, err)
		}
	}

	return &multiError{errors: errs}
}

// Combine combines the passed errors into a single error.
//
// If zero arguments were passed or if all items are nil, a nil error is
// returned.
//
// If only a single error was passed, it is returned as-is.
//
// Combine omits nil errors so this function may be used to combine
// together errors from operations that fail independently of each other.
//
// 	errors.Combine(
// 		reader.Close(),
// 		writer.Close(),
// 		pipe.Close(),
// 	)
//
// If any of the passed errors is already an aggregated error, it will be flattened along
// with the other errors.
//
// 	errors.Combine(errors.Combine(err1, err2), err3)
// 	// is the same as
// 	errors.Combine(err1, err2, err3)
//
// The returned error formats into a readable multi-line error message if
// formatted with %+v.
//
// 	fmt.Sprintf("%+v", multierr.Combine(err1, err2))
func Combine(errors ...error) error {
	return createMultiError(errors)
}
