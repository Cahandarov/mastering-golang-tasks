package user_inputs

import "fmt"

func Task5() {
	fmt.Println("Task-5   ******************")

	var age int
	fmt.Print("Enter your age:")
	_, err := fmt.Scan(&age)
	if err != nil {
		fmt.Println(err)
	}
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You can not pass")
	}
}
