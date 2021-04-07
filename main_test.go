package unwrapper

import (
	"fmt"
	"testing"
)

type CustomError struct {
	Msg string
}

func (e *CustomError) Error() string { return e.Msg + ": custom error" }

func TestUnwrappError1(t *testing.T) {
	t.Log("TestUnwrappError1")
	error1 := &CustomError{Msg: "error message"}
	error2 := fmt.Errorf("Error 2, %w", error1)
	error3 := fmt.Errorf("Error 3, %w", error2)

	stackError := unwrappError(error3)
	fmt.Println(stackError[0])
	result := []string{"*fmt.wrapError", "*fmt.wrapError", "*unwrapper.CustomError"}

	if len(stackError) != 3 {
		t.Fatal("Wrong len")
	}

	if result[0] != stackError[0] || result[1] != stackError[1] || result[2] != stackError[2] {
		t.Fatal("Wrong result")
	}
}
