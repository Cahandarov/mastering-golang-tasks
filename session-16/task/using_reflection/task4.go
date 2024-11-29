package using_reflection

import (
	"fmt"
	"reflect"
)

type Person2 struct {
	Name string
	Age  int
}

func (p Person2) Greet() string {
	return fmt.Sprintf("Hello, I am Alice.")
}

func InvokeMethod(value interface{}, methodName string) {
	val := reflect.ValueOf(value)
	kind := val.Kind()

	if kind == reflect.Struct {
		method := val.MethodByName(methodName)

		if method.IsValid() {
			fmt.Println("Invoked method:", methodName, "Result:", method.Call(nil)[0])
		} else {
			fmt.Println("Method Not Found:", methodName)
		}
	}

}
func Task4() {
	fmt.Println("Task 4   *******************")

	p := Person2{Name: "Ali", Age: 30}
	InvokeMethod(p, "Greet")
}
