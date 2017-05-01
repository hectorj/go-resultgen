# github.com/hectorj/go-resultgen

[![Build Status](https://travis-ci.org/hectorj/go-resultgen.svg?branch=master)](https://travis-ci.org/hectorj/go-resultgen)

A `go:generate` tool to generate some kind of [result types](https://en.wikipedia.org/wiki/Result_type) (except it isn't generic nor monadic).

The generated type exposes this interface:

```go
type MyTypeResult interface {
	// GetError tells us if the result is an error or not (in which case it returns `nil`).
	GetError() error
	// GetMyType gives us the encapsulated result value. Panics if the result is actually an error.
	GetMyType() MyType
}
```

It helps ensure a type cannot be in an invalid state, and is semantically more correct than a double return.
I also think it is nicer to read, but that is very subjective.

The traditional Go way:
```go
// you don't want to return something AND an error.
// you actually want to return something OR an error.
func buildSomething(someParam interface{}) (something, error) {
    if someParam == nil {
        return something{}, errors.New("someParam should not be nil")
    }
    return something{
        param: someParam,
    }, nil
}

func main() {
    smthing, err := buildSomething(nil)
    if err != nil {
        // We did check the error, so the errcheck linter won't complain
        log.Println(err)
    }
    // But we did not return! We still have our `something`, but in an invalid state.
    smthing.SomeMethod()
}
```

With a Result type:
```go
func buildSomething(someParam interface{}) somethingResult {
    if someParam == nil {
        return NewFailedSomethingResult(errors.New("someParam should not be nil"))
    }
    return NewValidSomethingResult(something{
        param: someParam,
    })
}

func main() {
    smthingResult := buildSomething(nil)
    if err := smthingResult.GetError(); err != nil {
        log.Println(err)
    }
    // Here, if there was an error, we will immediately panic.
    // This way, at no point in time did the user have an invalid `something`.
    smthing := smthingResult.GetSomething()
    smthing.SomeMethod()
}
```

## Installation

`go get -u github.com/hectorj/go-resultgen`

## Usage

It is meant to be used with the [`go generate`](https://blog.golang.org/generate) command.

See [the example](./tests/dummy_example_test.go).

There is also a "strict" mode, that will panic immediately if you try to get the value without checking the error at least once. I recommend it when running your tests. See [this other example](./tests/dummy_strict_example_test.go).
