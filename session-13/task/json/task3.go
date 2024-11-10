package json

import (
	"encoding/json"
	"fmt"
	"os"
)

type Students struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade int    `json:"grade"`
}

var sliceOfStudents = []Students{
	{Name: "Ali", Age: 20, Grade: 47},
	{Name: "Sabina", Age: 22, Grade: 61},
	{Name: "Medina", Age: 21, Grade: 81},
	{Name: "John", Age: 19, Grade: 45},
}

var students []Students

func createFileAndWriteData() {
	file, err := os.OpenFile("task/json/students.json", os.O_CREATE|os.O_RDWR, 0777)

	if err != nil {
		fmt.Println("Error happened on creating file:", err)
	}
	defer file.Close()

	j, err := json.Marshal(sliceOfStudents)
	if err != nil {
		fmt.Println("Error happened on converting to json:", err)
	}

	_, err = file.Write(j)
	if err != nil {
		fmt.Println("Error happened on writing to file:", err)
	}
}

func readFile() {
	var filteredStudents []Students
	data, err := os.ReadFile("task/json/students.json")
	if err != nil {
		fmt.Println("Error happened on reading file:", err)
	}
	err = json.Unmarshal(data, &students)
	if err != nil {
		fmt.Println("Error happened on unmarshalling json:", err)
	}
	for _, student := range students {
		if student.Grade >= 60 {
			filteredStudents = append(filteredStudents, student)
		}
	}
	if (len(filteredStudents)) > 0 {
		fmt.Println("Students with passing grades:")
		for _, student := range filteredStudents {
			fmt.Printf("- %s\n", student.Name)
		}
	}
}

func Task3() {
	fmt.Println("Task-3   ********************")
	createFileAndWriteData()
	readFile()
}
