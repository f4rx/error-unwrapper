package unwrapper

import (
	"fmt"
)

func unwrappError(err error) (stackError []string) {
	if err != nil {
		tr := true
		e := err
		fmt.Printf("Error:: %s\n", e)
		for tr {
			switch wrapped := e.(type) { //nolint:errorlint
			case interface{ Unwrap() error }:
				// fmt.Println("1", wrapped.Unwrap())
				// fmt.Printf("1: %T\n", wrapped)
				stackError = append(stackError, fmt.Sprintf("%T", wrapped.Unwrap()))
				fmt.Printf("1:: (%T): %v\n", wrapped.Unwrap(), wrapped.Unwrap())
				// t.Logf("1:: %T\n", wrapped.Unwrap())
				e = wrapped.Unwrap()
			case interface{ Cause() error }:
				fmt.Println("2", wrapped.Cause())
				tr = false
			default:
				stackError = append(stackError, fmt.Sprintf("%T", e))
				fmt.Printf("3:: (%T) %v\n", e, e)
				// fmt.Println("ololo:", e)
				tr = false
			}
		}
		// fmt.Printf("Error: %s, type: %T\n", err, err)
		// fmt.Printf("Error: %s, type: %T, %+v\n", err, err, err)
		// return
		// t.Fatal()
	}
	return
}
