package defining

import "fmt"

type Speaker interface {
	speak() string
}
type Dog struct {
	name string
}
type Person struct {
	name string
}

func (d Dog) speak() string {
	return fmt.Sprintf("Woof!")
}
func (p Person) speak() string {
	return fmt.Sprintf("Hello!")
}
func Task2() {
	fmt.Println("Task 2  ****************")

	dog := Dog{name: "Dogy"}
	person := Person{name: "John"}
	fmt.Printf("Dog says: %s\n", dog.speak())
	fmt.Printf("Person says: %s\n", person.speak())
}
