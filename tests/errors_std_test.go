// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors_test

import (
	"emperror.dev/errors"
	
	"testing"
)

func TestNewEqual(t *testing.T) {
	// Different allocations should not be equal.
	if errors.NewPlain("abc") == errors.NewPlain("abc") {
		t.Errorf(`New("abc") == New("abc")`)
	}
	if errors.NewPlain("abc") == errors.NewPlain("xyz") {
		t.Errorf(`New("abc") == New("xyz")`)
	}

	// Same allocation should be equal to itself (not crash).
	err := errors.NewPlain("jkl")
	if err != err {
		t.Errorf(`err != err`)
	}
}

func TestErrorMethod(t *testing.T) {
	err := errors.NewPlain("abc")
	if err.Error() != "abc" {
		t.Errorf(`New("abc").Error() = %q, want %q`, err.Error(), "abc")
	}
}

