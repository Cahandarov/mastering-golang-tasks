package methods

import (
	"fmt"
)

type Student struct {
	Name   string
	Grades []int
}

func (s Student) Average() int {
	var totalScore int = 0
	for _, v := range s.Grades {
		totalScore += v
	}
	return totalScore / len(s.Grades)
}

func Task4() {
	student1 := Student{
		Name:   "Ali",
		Grades: []int{80, 90, 70, 75, 90, 60},
	}

	student2 := Student{
		Name:   "Vali",
		Grades: []int{80, 90, 70, 75, 90, 60},
	}

	fmt.Printf("Student 1: %s, Average Grade: %.0d\n", student1.Name, student1.Average())
	fmt.Printf("Student 2: %s, Average Grade: %.0d\n", student2.Name, student2.Average())

	switch {
	case student1.Average() > student2.Average():
		fmt.Printf("%s has a higher average grade.", student1.Name)
	case student2.Average() > student1.Average():
		fmt.Printf("%s has a higher average grade.", student2.Name)
	default:
		fmt.Println("Average grades are equal.")
	}

}
