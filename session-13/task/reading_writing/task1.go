package reading_writing

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	name  string
	age   int
	grade int
}

func (s *Student) updateStudent(name string, age int, grade int) string {
	s.name = name
	s.age = age
	s.grade = grade
	return fmt.Sprintf("%s,%d,%d\n", s.name, s.age, s.grade)
}

var passingStudents []string

func updateCsvHeader() {
	file, err := os.OpenFile("task/reading_writing/data.csv", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("Error occurred:", err)
	}
	defer file.Close()

	r := bufio.NewReader(file)
	read, err := r.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Println("Error occurred:", err)
	}
	if !strings.Contains(read, "name,age,grade") {
		_, err = file.WriteString("name,age,grade\n")
		if err != nil {
			fmt.Println("Error occurred:", err)
		}
	}

}

func updateStudentsData(str string) {
	file, err := os.OpenFile("task/reading_writing/data.csv", os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Error occurred:", err)
	}
	defer file.Close()

	_, err = file.WriteString(str)
	if err != nil {
		fmt.Println("Error occurred:", err)
	}
}

func listStudents() {
	file, err := os.OpenFile("task/reading_writing/data.csv", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error occurred:", err)
	}
	r := bufio.NewReader(file)

	_, err = r.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading the header:", err)
		return
	}

	for {
		read, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error occurred while reading:", err)
			return
		}
		str := strings.Split(strings.TrimSpace(read), ",")
		grade, err := strconv.Atoi(str[2])
		if err != nil {
			fmt.Println("Error occurred while converting grade:", err)
		}
		if grade > 60 {
			passingStudents = append(passingStudents, str[0])
		}
	}

	if len(passingStudents) > 0 {
		fmt.Println("Students with passing grades:")
		for _, student := range passingStudents {
			fmt.Printf("- %s\n", student)
		}
	} else {
		fmt.Println("No students with passing grades.")
	}
}

func Task1() {
	fmt.Println("Task-1   ********************")

	updateCsvHeader()
	Alice := Student{}
	updateStudentsData(Alice.updateStudent("Alice", 20, 85))
	Bob := Student{}
	updateStudentsData(Bob.updateStudent("Bob", 22, 55))
	Carol := Student{}
	updateStudentsData(Carol.updateStudent("Carol", 21, 70))

	listStudents()

}
