package custerr

import (
	"fmt"
)

type DivisionError struct {
	Message string
}

var divisionError = DivisionError{
	Message: "Division by zero is not allowed.",
}

func (e *DivisionError) Error() string {
	return fmt.Sprintf("Error: %s", e.Message)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, &divisionError
	} else {
		return a / b, nil
	}
}

func Task1() {

	fmt.Println("Task 1 - *****************************")
	var res, err = divide(4, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result: ", res)
	}

}
