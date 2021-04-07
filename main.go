package unwrapper

import (
	"fmt"
)

func unwrappError(err error) (stackError []string) {
	if err != nil {
		tr := true
		e := err
		// fmt.Println(e)
		stackError = append(stackError, fmt.Sprintf("%T", e))

		for tr {
			switch wrapped := e.(type) { //nolint:errorlint
			case interface{ Unwrap() error }:
				stackError = append(stackError, fmt.Sprintf("%T", wrapped.Unwrap()))
				e = wrapped.Unwrap()
				// fmt.Println(e)
			case interface{ Cause() error }:
				fmt.Println("2", wrapped.Cause())
				tr = false
			default:
				tr = false
			}
		}
	}
	return
}
