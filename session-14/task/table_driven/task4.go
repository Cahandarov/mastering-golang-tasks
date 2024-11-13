package table_driven

import (
	"errors"
	"math"
)

var ErrorWrongInput = errors.New("wrong input entered")

func ConvertTemperature(value float64, scale string) (float64, error) {

	if scale == "Celsius" {
		return math.Round(((value * 1.8) + 32)), nil
	} else if scale == "Fahrenheit" {
		return math.Round(((value - 32) * (5.0 / 9.0))), nil
	} else {
		return 0, ErrorWrongInput
	}
}

//func Task4() {
//	fmt.Println("Task-4    **************")
//
//	fmt.Println(ConvertTemperature(28, "Celsius"))
//	fmt.Println(ConvertTemperature(33.8, "Fahrenheit"))
//	fmt.Println(ConvertTemperature(56, "Fahrenheit"))
//	fmt.Println(ConvertTemperature(35, "Fahrenheitttt"))
//}
