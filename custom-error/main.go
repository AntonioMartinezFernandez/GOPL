package main

import (
	"fmt"
	"math/rand"
)

/*
Error Types
*/
type LengthError struct {
	msg string
}

func (e *LengthError) Error() string { return e.msg }

type ValueError struct {
	msg string
}

func (e *ValueError) Error() string { return e.msg }

/*
Returns 'LengthError' or 'ValueError' randomly
*/
func giveMeAnError() interface{} {
	r := rand.Intn(2)

	lengthError := LengthError{msg: "this is a length error message"}
	valueError := ValueError{msg: "this is a value error message"}

	switch r {
	case 0:
		return lengthError
	case 1:
		return valueError
	}

	return nil
}

/*
Give me 5 (errors)
*/
func main() {
	for i := 0; i < 5; i++ {
		err := giveMeAnError()

		// LengthError Type Casting
		le, ok := err.(LengthError)
		if ok {
			fmt.Printf("🔥 LengthError received\n 📜 %s\n\n", le.Error())
		}

		// ValueError Type Casting
		ve, ok := err.(ValueError)
		if ok {
			fmt.Printf("🔥 ValueError received\n 📜 %s\n\n", ve.Error())
		}
	}
}
