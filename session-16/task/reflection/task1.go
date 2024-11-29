package reflection

import (
	"fmt"
	"reflect"
)

func IdentifyTypeAndKind(value interface{}) {
	typ := reflect.TypeOf(value)
	val := reflect.ValueOf(value)
	kind := val.Kind()
	fmt.Println("Value:", val, "Type:", typ, "Kind:", kind)

}

func Task1() {
	fmt.Println("Task 1   *******************")
	IdentifyTypeAndKind(42)
	IdentifyTypeAndKind("Hello")
	IdentifyTypeAndKind([]int{1, 2, 3})
}
