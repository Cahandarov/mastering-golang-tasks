package embedding

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}
type Employee struct {
	Person
	EmployeeID int
	Position   string
}

func Task5() {
	employee1 := Employee{
		Person: Person{
			FirstName: "Ali",
			LastName:  "Aliyev",
		},
		EmployeeID: 12345,
		Position:   "Software Engineer",
	}
	fmt.Printf("Name: %s %s\nEmployee ID: %d\nPosition: %s\n", employee1.FirstName, employee1.LastName, employee1.EmployeeID, employee1.Position)

}
