package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func session2() {
	var ProjectStatus string = "in progress"
	const NumberOfTotalTasks int = 100
	CompletedTasks := 25
	const (
		Low = iota
		Medium
		High
	)
	var ProjectCompletionStatus bool = true
	if ProjectStatus == "in progress" {
		ProjectCompletionStatus = false
	} else {
		ProjectCompletionStatus = true
	}
	fmt.Println("Current project status:", strings.ToUpper(ProjectStatus))
	fmt.Printf("Tasks completed: %d out of %d\n", CompletedTasks, NumberOfTotalTasks)
	fmt.Printf("Task priorities: %d-Low, %d-Medium, %d-High\n", Low, Medium, High)
	fmt.Printf("Is the project completed? %s", strconv.FormatBool(ProjectCompletionStatus))

}
func main() {
	// First print
	fmt.Println("Welcome to the Task Management System!")

	// Second print
	fmt.Println("Project start date is:", time.Date(2024, time.September, 18, 0, 0, 0, 0, time.UTC))

	// Third print
	fmt.Println("Project: Task Management System")
	fmt.Println()
	session2()
}

//I made changes in task
//I made second changes in task
