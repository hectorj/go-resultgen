// +build !strict

package tests_test

import (
	"fmt"

	"github.com/hectorj/go-resultgen/tests"
	"github.com/pkg/errors"
)

/*
// In the tests package, we have:
//go:generate go run ../main.go Dummy --tags=!strict
type Dummy struct {
	ID int
}
*/

func Example() {
	// We get a result. We don't know yet if we have an error, or a valid Dummy instance.
	result := DummyGetter(true)

	// So we check for an error first.
	if err := result.GetError(); err != nil {
		// Here is our error processing code.
		// In real life you would probably use a more sensible logging strategy, or just return the error.
		// The important point is that we won't call result.GetDummy() if there is an error.
		fmt.Println(1, "error:", err)
		return
	}

	// As you will see in the ouput, there is no error
	fmt.Println(1, "id:", result.GetDummy().ID)

	// Let's try again
	result2 := DummyGetter(false)
	if err := result2.GetError(); err != nil {
		// As you will see in the ouput, this time we actually have an error
		fmt.Println(2, "error:", err)
	} else {
		fmt.Println(2, "id:", result2.GetDummy().ID)
		return
	}

	// The following examples are unsafe, they may panic.
	defer func() {
		if panicErr := recover(); panicErr != nil {
			fmt.Println("panic:", panicErr)
		}
	}()

	result3 := DummyGetter(true)
	// No error check, YOLO
	fmt.Println(3, "id:", result3.GetDummy().ID) // Does not panic because the result is valid

	result4 := DummyGetter(false)
	// Playing russian roulette here
	fmt.Println(4, "id:", result4.GetDummy().ID) // Panics. We played, we lost.

	// Output:
	// 1 id: 42
	// 2 error: invalid
	// 3 id: 42
	// panic: invalid
}

func DummyGetter(valid bool) tests.DummyResult {
	if valid {
		return tests.NewValidDummyResult(tests.Dummy{
			ID: 42,
		})
	}

	return tests.NewFailedDummyResult(errors.New("invalid"))
}
