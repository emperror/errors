package match

import (
	"errors"
	"testing"
)

func TestAny(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		matcher := Any{}

		if matcher.MatchError(errors.New("error")) {
			t.Error("empty any matcher is not supposed to match an error")
		}
	})

	t.Run("not_empty", func(t *testing.T) {
		matcher := Any{
			errorMatcherFunc(func(err error) bool { return false }),
			errorMatcherFunc(func(err error) bool { return true }),
		}

		if !matcher.MatchError(errors.New("error")) {
			t.Error("not-empty any matcher is supposed to match an error if any of the matchers match it")
		}
	})
}

func TestAll(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		matcher := All{}

		if !matcher.MatchError(errors.New("error")) {
			t.Error("empty all matcher is not supposed to match an error")
		}
	})

	t.Run("not_empty", func(t *testing.T) {
		matcher := All{
			errorMatcherFunc(func(err error) bool { return false }),
			errorMatcherFunc(func(err error) bool { return true }),
		}

		if matcher.MatchError(errors.New("error")) {
			t.Error("not-empty all matcher is not supposed to match an error if not all of the matchers match it")
		}
	})
}

func TestIs(t *testing.T) {
	err := errors.New("error")

	matcher := Is(err)

	if !matcher.MatchError(err) {
		t.Error("is matcher is supposed to match an error if errors.Is returns true")
	}
}

type asErrorStub struct{}

func (asErrorStub) Error() string {
	return "error"
}

func (asErrorStub) IsError() bool {
	return true
}

func TestAs(t *testing.T) {
	var matchErr interface {
		IsError() bool
	}

	matcher := As(&matchErr)

	if !matcher.MatchError(asErrorStub{}) {
		t.Error("is matcher is supposed to match an error if errors.Is returns true")
	}
}
