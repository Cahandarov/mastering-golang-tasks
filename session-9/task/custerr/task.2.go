package custerr

import (
	"fmt"
)

type EnhancedDivisionErrors struct {
	Dividend float64
	Divisor  float64
	DivisionError
}

func (e *EnhancedDivisionErrors) Error() string {
	return fmt.Sprintf("Cannot divide %v by %v.", e.Dividend, e.Divisor)
}

func divide2(a, b float64) (float64, error) {
	if b == 0 {
		return 0, &EnhancedDivisionErrors{
			Dividend: a,
			Divisor:  b,
		}
	} else {
		return a / b, nil
	}
}

func Task2() {

	fmt.Println("Task 2 - *****************************")
	var res, err = divide2(4, 0)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println("Result: \n", res)
	}

}
