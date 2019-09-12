# Can

[![GoDoc](https://godoc.org/github.com/kellengreen/can?status.svg)](https://godoc.org/github.com/kellengreen/can)

The `can` package helps bring simplified error handling to go.
Works similar to `try`/`catch` from other languages, while utilizing go's `panic`, `defer`, and `recover` statements alongside named return values.


# Usage
```go
// fail
fmt.Println(func() (_ string, e error) {
    defer can.Recover(&e)
    can.Panic(errors.New("OOPS"))
    return "OK", nil
}())

// succeed
fmt.Println(func() (_ string, e error) {
    defer can.Recover(&e)
    can.Panic(nil)
    return "OK", nil
}())

// Output:
// OOPS
// OK <nil>
```