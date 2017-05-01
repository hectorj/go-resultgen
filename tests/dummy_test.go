// +build !strict

package tests

import (
	"errors"
	"testing"
)

func TestNewValidDummyResult_GetDummy(t *testing.T) {
	result := NewValidDummyResult(Dummy{
		ID: 42,
	})

	if expected, actual := 42, result.GetDummy().ID; expected != actual {
		t.Fatalf("expected %d, got %d (value seems to have changed when encapsulated, while it shouldn't)", expected, actual)
	}
}

func TestNewValidDummyResult_GetError(t *testing.T) {
	result := NewValidDummyResult(Dummy{
		ID: 42,
	})

	if actual := result.GetError(); nil != actual {
		t.Fatalf("expected no error, got %q", actual)
	}
}

func TestNewFailedDummyResult_nil(t *testing.T) {
	var panicErr interface{}
	func() {
		defer func() {
			panicErr = recover()
		}()
		NewFailedDummyResult(nil)
	}()

	if panicErr == nil {
		t.Fatal("expected a panic")
	}
}

func TestNewFailedDummyResult_GetError(t *testing.T) {
	expected := errors.New("expected error")
	result := NewFailedDummyResult(expected)

	if actual := result.GetError(); expected != actual {
		t.Fatalf("expected %q, got %+v", expected, actual)
	}
}

func TestNewFailedDummyResult_GetDummy(t *testing.T) {
	expected := errors.New("expected error")
	result := NewFailedDummyResult(expected)

	var panicErr interface{}
	func() {
		defer func() {
			panicErr = recover()
		}()
		result.GetDummy()
	}()

	if panicErr == nil {
		t.Fatal("expected a panic")
	}
}
