package can

import (
	"errors"
	"fmt"
	"testing"
)

func Example() {
	// fail
	fmt.Println(func() (_ string, e error) {
		defer Recover(&e)
		Panic(errors.New("OOPS"))
		return "OK", nil
	}())

	// succeed
	fmt.Println(func() (_ string, e error) {
		defer Recover(&e)
		Panic(nil)
		return "OK", nil
	}())

	// Output:
	// OOPS
	// OK <nil>
}

func TestPanicNil(t *testing.T) {
	defer func() {
		if recover() != nil {
			t.Fatal("recover was not nil")
		}
	}()
	Panic(nil)
}

func TestPanicErr(t *testing.T) {

	p := errors.New("OOPS")

	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("recover was nil")
		}

		e, ok := r.(error)
		if ok == false {
			t.Fatal("recover was not an error interface")
		}

		if e.Error() != p.Error() {
			t.Fatal("recover error message did not match")
		}
	}()

	Panic(p)
	t.Fatal("panic was not called")
}

func TestRecoverNil(t *testing.T) {
	var e error
	Recover(&e)
	if e != nil {
		t.Fatal("error was modified")
	}
}

func TestRecoverErr(t *testing.T) {
	p := errors.New("p")

	f := func() (e error) {
		defer Recover(&e)
		panic(p)
	}

	r := f()
	if p.Error() != r.Error() {
		t.Fatal("error did not match")
	}
}

func TestRecoverStr(t *testing.T) {
	p := "ERR"

	f := func() (e error) {
		defer Recover(&e)
		panic(p)
	}

	r := f()
	if p != r.Error() {
		t.Fatal("error did not match")
	}
}
