package match

import (
	"emperror.dev/errors"
)

// ErrorMatcher checks if an error matches a certain condition.
type ErrorMatcher interface {
	// MatchError checks if err matches a certain condition.
	MatchError(err error) bool
}

// errorMatcherFunc turns a plain function into an ErrorMatcher if it's definition matches the interface.
type errorMatcherFunc func(err error) bool

// MatchError calls the underlying function to check if err matches a certain condition.
func (fn errorMatcherFunc) MatchError(err error) bool {
	return fn(err)
}

// Any matches an error if any of the underlying matchers match it.
type Any []ErrorMatcher

// MatchError calls underlying matchers with err.
// If any of them matches err it returns true, otherwise false.
func (m Any) MatchError(err error) bool {
	for _, matcher := range m {
		if matcher.MatchError(err) {
			return true
		}
	}

	return false
}

// All matches an error if all of the underlying matchers match it.
type All []ErrorMatcher

// MatchError calls underlying matchers with err.
// If all of them matches err it returns true, otherwise false.
func (m All) MatchError(err error) bool {
	for _, matcher := range m {
		if !matcher.MatchError(err) {
			return false
		}
	}

	return true
}

// Is returns an error matcher that determines matching by calling errors.Is.
func Is(target error) ErrorMatcher {
	return errorMatcherFunc(func(err error) bool {
		return errors.Is(err, target)
	})
}

// As returns an error matcher that determines matching by calling errors.As.
func As(target interface{}) ErrorMatcher {
	return errorMatcherFunc(func(err error) bool {
		// TODO recover from panic in case target is invalid to avoid and error storm
		return errors.As(err, target)
	})
}
