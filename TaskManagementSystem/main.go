package main

import (
	"fmt"
	"strconv"
)

const NumberOfTotalTasks int = 100

// #1 Setting project status depend on completed task
func taskCompletionCheck(CompletedTasks int) string {
	var ProjectStatus string
	if CompletedTasks < NumberOfTotalTasks {
		ProjectStatus = "IN PROGRESS"
	} else {
		ProjectStatus = "FINISHED"
	}
	return ProjectStatus
}

// #2 Remaining tasks calculation depends on completed tasks and returning multiple returns
func taskRemainingCalculation(CompletedTasks int) (int, bool) {
	var nearCompletion bool
	RemainingTasks := NumberOfTotalTasks - CompletedTasks
	if CompletedTasks > 90 {
		nearCompletion = true
	} else {
		nearCompletion = false
	}
	return RemainingTasks, nearCompletion
}

// #3 Listing completed tasks
func completedTaskListing(CompletedTasks int) {
	if CompletedTasks > 0 {
		fmt.Println("Task list:")
		for t := 1; t <= CompletedTasks; t++ {
			fmt.Printf("%d: Task %d\n", t, t)
		}
	}
}

// #4 Listing countdown of remaining tasks by using recursive function
func recursiveCountdownOfRemainingTasks(RemainingTasks int) {
	if RemainingTasks > 0 {
		fmt.Printf("Tasks remaining : %d\n", RemainingTasks)
		recursiveCountdownOfRemainingTasks(RemainingTasks - 1)
	} else if RemainingTasks == 0 {
		fmt.Println("Countdown Of Remaining Tasks Completed!")
	}

}

// #5 Setting project phases depend on completed tasks and using panic when default
func projectPhases(CompletedTasks int) string {
	var ProjectPhase string
	switch {
	case CompletedTasks > -1 && CompletedTasks < 30:
		ProjectPhase = "Project is in the starting phase"
	case CompletedTasks >= 30 && CompletedTasks <= 60:
		ProjectPhase = "Project is in the midway"
	case CompletedTasks > 60 && CompletedTasks < 100:
		ProjectPhase = "Project is in the final phase"
	case CompletedTasks == 100:
		ProjectPhase = "Project is completed"
	default:
		panic("tasksCompleted exceeds total tasks")
	}
	return ProjectPhase
}

// #6 Checking project status by checking project status
func taskStatus(ProjectStatus string) bool {
	var ProjectCompletionStatus bool
	if ProjectStatus == "IN PROGRESS" {
		ProjectCompletionStatus = false
	} else {
		ProjectCompletionStatus = true
	}
	return ProjectCompletionStatus
}

// #7 Adding new tasks by using variadic function
func variadicTaskLists(tasks ...string) {
	for key, task := range tasks {
		fmt.Printf("%d. %s\n", key+1, task)
	}
}

// #8 I am using task update function to pass completed tasks
func taskUpdate(CompletedTasks int) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic: Error:", r)
		}
	}()

	RemainingTasks, nearCompletion := taskRemainingCalculation(CompletedTasks) //#2

	if nearCompletion {
		fmt.Println("Project is near completion!")
	}

	ProjectPhase := projectPhases(CompletedTasks) //#5
	fmt.Printf("%s\n", ProjectPhase)

	completedTaskListing(CompletedTasks) //#3

	if RemainingTasks > 0 {
		fmt.Println("Recursive countdown of remaining tasks:")
		recursiveCountdownOfRemainingTasks(RemainingTasks) //#4
	}

	ProjectStatus := taskCompletionCheck(CompletedTasks) //#1

	ProjectCompletionStatus := taskStatus(ProjectStatus) //#6
	fmt.Printf("Is the project completed? %s\n", strconv.FormatBool(ProjectCompletionStatus))
}

func main() {

	taskUpdate(100) //#8

	fmt.Println("Received new tasks:")
	variadicTaskLists("Session-1 task", "Session-2 task", "Session-3 task") //#7
}
