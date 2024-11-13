package switching

import "fmt"

type TypeSwitch interface{}

func CheckType(content any) {
	switch v := content.(type) {
	case int:
		fmt.Printf("Type is int: %d\n", v)
	case string:
		fmt.Printf("Type is string: %s\n", v)
	case bool:
		fmt.Printf("Type is bool: %t\n", v)
	default:
		fmt.Println("Unknown type")
	}

}
func Task6() {
	fmt.Println("Task 6  ****************")
	content1 := 100
	content2 := "Hi Gophers"
	content3 := true
	content4 := 67.4
	CheckType(content1)
	CheckType(content2)
	CheckType(content3)
	CheckType(content4)
}
