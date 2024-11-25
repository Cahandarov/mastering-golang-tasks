package using_reflection

import (
	"errors"
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
	City string
}

var ErrorWrongInput = errors.New("value must be a pointer to a struct")
var ErrorWrongType = errors.New("wrong entered data type")

func SetFieldValue(value interface{}, fieldName string, newValue interface{}) error {
	val := reflect.ValueOf(value)
	//for update structs fields we need to receive pointer so we check reflect.Ptr
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return ErrorWrongInput
	}
	str := val.Elem()
	enteredFiled := str.FieldByName(fieldName)
	if !enteredFiled.IsValid() {
		return fmt.Errorf("%s field not exists", fieldName)
	}
	newVal := reflect.ValueOf(newValue)
	if newVal.Type() != enteredFiled.Type() {
		return ErrorWrongType
	}
	enteredFiled.Set(newVal)

	return nil
}

func Task3() {
	fmt.Println("Task 3   *******************")

	p := &Person{Name: "Ali", Age: 30, City: "Baku"}
	err := SetFieldValue(p, "City", "Ganja")
	fmt.Println(err)
	fmt.Println(p)
}
