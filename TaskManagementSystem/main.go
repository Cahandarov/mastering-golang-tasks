package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func TaskUpdate(CompletedTasks int) {

	const NumberOfTotalTasks int = 100
	var ProjectStatus string

	RemainingTasks := NumberOfTotalTasks - CompletedTasks
	if CompletedTasks != NumberOfTotalTasks {
		ProjectStatus = "in progress"
	} else {
		ProjectStatus = "finished"
	}

	var ProjectCompletionStatus bool
	if ProjectStatus == "in progress" {
		ProjectCompletionStatus = false
	} else {
		ProjectCompletionStatus = true
	}

	const (
		Low = iota + 1
		Medium
		High
	)

	var ProjectPhase string
	switch {
	case CompletedTasks < 30:
		ProjectPhase = "Starting phase"
	case 30 <= CompletedTasks && CompletedTasks <= 60:
		ProjectPhase = "Midway"
	case 60 <= CompletedTasks:
		ProjectPhase = "Final phase"
	default:
		fmt.Println("Wrong entered data")
	}

	fmt.Printf("Tasks remaining %d out of %d\n", RemainingTasks, NumberOfTotalTasks)
	fmt.Printf("Current project status: %s\n", strings.ToUpper(ProjectStatus))
	fmt.Printf("Project is in the %s.\n", strings.ToLower(ProjectPhase))
	fmt.Printf("Task priorities: %d-Low, %d-Medium, %d-High\n", Low, Medium, High)
	if CompletedTasks > 0 {
		fmt.Println("Task list:")
		for t := 1; t <= RemainingTasks; t++ {
			fmt.Printf("- Task %d\n", t)
		}
	}

	fmt.Printf("Is the project completed? %s\n", strconv.FormatBool(ProjectCompletionStatus))

}

func main() {
	// First print
	fmt.Println("Welcome to the Task Management System!")

	// Second print
	fmt.Println("Project start date is:", time.Date(2024, time.September, 18, 0, 0, 0, 0, time.UTC))

	// Third print
	fmt.Println("Project: Task Management System")
	fmt.Println()
	TaskUpdate(95)
	fmt.Println()
	fmt.Println()
}
