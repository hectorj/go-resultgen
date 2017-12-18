// +build !strict

package tests

import (
	"errors"
)

// File generated by github.com/hectorj/go-resultgen.

// DummyResult is a result type for Dummy.
// See https://en.wikipedia.org/wiki/Result_type .
type DummyResult struct {
	value Dummy
	err error
}

func NewValidDummyResult(value Dummy) DummyResult {
	return DummyResult{
		value: value,
		err: nil,
	}
}

func NewFailedDummyResult(err error) DummyResult {
	if err == nil {
		panic(errors.New("cannot create failed result from nil error"))
	}
	return DummyResult{
		err: err,
	}
}

func (r DummyResult) isDummyResult() {}

func (r DummyResult) GetError() error {
	return r.err
}

func (r DummyResult) GetDummy() Dummy {
	if r.err != nil {
		panic(r.err)
	}
	return r.value
}
