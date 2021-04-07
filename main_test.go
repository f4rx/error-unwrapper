package unwrapper

import (
	"testing"
	// "errors"
	"fmt"
)

type CustomError struct {
    Msg string
}

func (e *CustomError) Error() string { return e.Msg + ": custom error" }

func TestUnwrappError1(t *testing.T) {

	error1 := &CustomError{Msg: "error message"}
	error2 := fmt.Errorf("Error 2, %w", error1)
	error3 := fmt.Errorf("Error 3, %w", error2)


	// if err != nil {
	// 	fmt.Println("error:", err)
	// }

	stackError := unwrappError(error3)
	fmt.Println(stackError)
}

func TestUnwrappError2(t *testing.T) {

	error1 := &CustomError{Msg: "error message"}
	// error2 := fmt.Errorf("Error 2, %w", error1)
	// error3 := fmt.Errorf("Error 3, %w", error2)


	// if err != nil {
	// 	fmt.Println("error:", err)
	// }

	stackError := unwrappError(error1)
	fmt.Println(stackError)
}
