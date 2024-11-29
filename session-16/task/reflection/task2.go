package reflection

import (
	"errors"
	"fmt"
	"reflect"
)

var ErrorWrongInput = errors.New("wrong input")

func InspectStruct(value interface{}) error {
	val := reflect.ValueOf(value)
	typ := reflect.TypeOf(value)
	kind := val.Kind()
	if kind != reflect.Struct {
		return ErrorWrongInput
	}

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fmt.Println("Filed Name:", field.Name, "Type:", field.Type, "Value:", val.Field(i))
	}
	return nil
}

func Task2() {
	fmt.Println("Task 2   *******************")

	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Ali", Age: 30}
	err := InspectStruct(p)
	if err != nil {
		fmt.Println(err)
	}
}
