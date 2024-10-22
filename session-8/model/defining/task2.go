package defining

import "fmt"

type Speaker interface {
	Speak() string
}
type Dog struct {
	Name string
}
type Person struct {
	Name string
}

func (d Dog) Speak() string {
	return fmt.Sprintf("Woof!")
}
func (p Person) Speak() string {
	return fmt.Sprintf("Hello!")
}

func Task2() {
	fmt.Println("Task 2  ****************")
	var dog Speaker
	dog = Dog{Name: "Dogy"}

	var person Speaker
	person = Person{Name: "John"}
	fmt.Printf("Dog says: %s\n", dog.Speak())
	fmt.Printf("Person says: %s\n", person.Speak())
}
