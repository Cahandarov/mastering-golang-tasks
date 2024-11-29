package xml

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

type Employee struct {
	Name     string `xml:"name"`
	Position string `xml:"position"`
	Salary   int64  `xml:"salary"`
}

type employees struct {
	Employees []Employee `xml:"employee"`
}

var employeesList = employees{}
var unMarshalledData = employees{}

func (e *Employee) updateEmployee(name string, position string, salary int64) Employee {
	e.Name = name
	e.Position = position
	e.Salary = salary
	return *e
}

func createAndUpdateXML(emp Employee) {

	employeesList.Employees = append(employeesList.Employees, emp)

	file, err := os.OpenFile("task/xml/employees.xml", os.O_CREATE|os.O_RDWR|os.O_APPEND|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println("Error occurred while creating file", err)
	}
	defer file.Close()

	x, err := xml.Marshal(employeesList)
	if err != nil {
		fmt.Println("Error occurred while marshalling", err)
	}

	_, err = file.Write(x)
	if err != nil {
		fmt.Println("Error occurred while writing", err)
	}
}

func readXMLFile() {
	var filteredSlice []string
	data, err := os.ReadFile("task/xml/employees.xml")
	if err != nil {
		fmt.Println("Error happened on reading file:", err)
	}
	err = xml.Unmarshal(data, &unMarshalledData)
	if err != nil {
		fmt.Println("Error happened on unmarshalling json:", err)
	}
	for _, v := range unMarshalledData.Employees {
		if v.Salary > 50000 {
			filteredSlice = append(filteredSlice, v.Name)
		}
	}
	fmt.Println("Employees with Salary above 50000:")
	fmt.Printf("- %s\n", strings.Join(filteredSlice, ", "))
}
func Task5() {
	fmt.Println("Task-5   ********************")
	emp1 := Employee{}
	emp2 := Employee{}
	createAndUpdateXML(emp1.updateEmployee("Natig", "Manager", 75000))
	createAndUpdateXML(emp2.updateEmployee("Subhan", "Developer", 45000))
	readXMLFile()
}

//<employees><employee><name>Natig</name><position>Manager</position><salary>75000</salary></employee><employee><name>Subhan</name><position>Developer</position><salary>45000</salary></employee></employees>
