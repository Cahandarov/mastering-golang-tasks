package unit_testing

import (
	"errors"
)

var ErrorZeroDivide = errors.New("divide by zero")
var ErrorUnknownInput = errors.New("unknown input entered")

func Calculate(a, b int, operation string) (int, error) {
	if operation == "+" {
		return a + b, nil
	} else if operation == "-" {
		return a - b, nil
	} else if operation == "*" {
		return a * b, nil
	} else if operation == "/" && b != 0 {
		return a / b, nil
	} else if operation == "/" && b == 0 {
		return 0, ErrorZeroDivide
	} else {
		return 0, ErrorUnknownInput
	}
}

//func Task1() {
//fmt.Println("Task-1    **************8")
//fmt.Println(Calculate(2, 5, "+"))
//fmt.Println(Calculate(5, 5, "-"))
//fmt.Println(Calculate(2, 5, "*"))
//fmt.Println(Calculate(10, 5, "/"))
//fmt.Println(Calculate(2, 0, "/"))
//fmt.Println(Calculate(2, 0, ")"))
//}
