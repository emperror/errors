package errors

import (
	"reflect"
	"testing"
)

func TestWithDetails(t *testing.T) {
	origErr := NewPlain("something went wrong")
	details := []interface{}{"key", "value"}
	err := WithDetails(origErr, details...)

	t.Parallel()

	t.Run("error_message", func(t *testing.T) {
		checkErrorMessage(t, err, "something went wrong")
	})

	t.Run("unwrap", func(t *testing.T) {
		checkUnwrap(t, err, origErr)
	})

	t.Run("format", func(t *testing.T) {
		checkFormat(t, err, map[string][]string{
			"%s":  {"something went wrong"},
			"%q":  {`"something went wrong"`},
			"%v":  {"something went wrong"},
			"%+v": {"something went wrong"},
		})
	})

	t.Run("nil", func(t *testing.T) {
		checkErrorNil(t, WithDetails(nil, "key", "value"))
	})

	t.Run("details", func(t *testing.T) {
		d := err.(*withDetails).Details()

		for i, detail := range d {
			if got, want := detail, details[i]; got != want {
				t.Errorf("error detail does not match the expected one\nactual:   %+v\nexpected: %+v", got, want)
			}
		}
	})

	t.Run("details_missing_value", func(t *testing.T) {
		details := []interface{}{"key", nil}
		err := WithDetails(origErr, "key")

		d := err.(*withDetails).Details()

		for i, detail := range d {
			if got, want := detail, details[i]; got != want {
				t.Errorf("error detail does not match the expected one\nactual:   %+v\nexpected: %+v", got, want)
			}
		}
	})
}

func TestGetDetails(t *testing.T) {
	err := WithDetails(
		WithMessage(
			WithDetails(
				Wrap(
					WithDetails(
						New("error"),
						"key", "value",
					),
					"wrapped error",
				),
				"key2", "value2",
			),
			"another wrapped error",
		),
		"key3", "value3",
	)

	expected := []interface{}{
		"key", "value",
		"key2", "value2",
		"key3", "value3",
	}

	actual := GetDetails(err)

	if got, want := actual, expected; !reflect.DeepEqual(got, want) {
		t.Errorf("context does not match the expected one\nactual:   %v\nexpected: %v", got, want)
	}
}
