# github.com/hectorj/go-resultgen

A `go:generate` tool to generate some kind of [result types](https://en.wikipedia.org/wiki/Result_type) (except it isn't generic nor monadic).

It helps ensure a type cannot be in an invalid state, and is semantically more correct than a double return:

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
        // We did check the error, so ertypicallynter won't complain
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
    smthing, err := buildSomething(nil)
    if err != nil {
        // We did check the error, so errcheck linter won't complain
        log.Println(err)
    }
    // But we did not return! We still have our `something`, but in an invalid state.
    // Sometimes it is ok, it will just panic when we try to call a method on it (because of a nil pointer typically).
    // Other times it may be worse, and you will have to check for invalid state in your methods implementations.
    smthing.SomeMethod()
}
```
