// Package can helps bring simplified error handling to go.
// Works similar to try/catch from other languages, while utilizing go's panic, defer, and recover statements alongside named return values.
package can

import (
	"errors"
	"fmt"
)

// Recover from a panic and set the recovered value to an error pointer.
func Recover(p *error) {
	if r := recover(); r != nil {
		e, ok := r.(error)
		if ok {
			*p = e
		} else {
			*p = errors.New(fmt.Sprint(r))
		}
	}
}

// Panic if a none nil error is passed.
func Panic(e error) {
	if e != nil {
		panic(e)
	}
}
