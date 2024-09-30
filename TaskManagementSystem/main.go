package main

import (
	"fmt"
	"strconv"
)

const NumberOfTotalTasks int = 100

func taskCompletionCheck(CompletedTasks int) string {
	var ProjectStatus string
	if CompletedTasks < NumberOfTotalTasks {
		ProjectStatus = "IN PROGRESS"
	} else {
		ProjectStatus = "FINISHED"
	}
	return ProjectStatus
}

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

func completedTaskListing(CompletedTasks int) {
	if CompletedTasks > 0 {
		fmt.Println("Task list:")
		for t := 1; t <= CompletedTasks; t++ {
			fmt.Printf("%d: Task %d\n", t, t)
		}
	}
}

func recursiveCountdownOfRemainingTasks(RemainingTasks int) {
	if RemainingTasks > 0 {
		fmt.Printf("Tasks remaining : %d\n", RemainingTasks)
		recursiveCountdownOfRemainingTasks(RemainingTasks - 1)
	} else if RemainingTasks == 0 {
		fmt.Println("Countdown Of Remaining Tasks Completed!")
	}

}
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

func taskStatus(ProjectStatus string) bool {
	var ProjectCompletionStatus bool
	if ProjectStatus == "IN PROGRESS" {
		ProjectCompletionStatus = false
	} else {
		ProjectCompletionStatus = true
	}
	return ProjectCompletionStatus
}

func variadicTaskLists(tasks ...string) {
	for key, task := range tasks {
		fmt.Printf("%d. %s\n", key+1, task)
	}
}

func taskUpdate(CompletedTasks int) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic: Error:", r)
		}
	}()

	ProjectPhase := projectPhases(CompletedTasks) //Project phase
	fmt.Printf("%s\n", ProjectPhase)

	completedTaskListing(CompletedTasks) //Completed tasks listing

	RemainingTasks, _ := taskRemainingCalculation(CompletedTasks) //RECURSIVE Remaining tasks listing
	if RemainingTasks > 0 {
		fmt.Println("Recursive countdown of remaining tasks:")
		recursiveCountdownOfRemainingTasks(RemainingTasks)
	}

	ProjectStatus := taskCompletionCheck(CompletedTasks)

	ProjectCompletionStatus := taskStatus(ProjectStatus)

	fmt.Printf("Is the project completed? %s\n", strconv.FormatBool(ProjectCompletionStatus))

}

func main() {

	taskUpdate(100)

	fmt.Println("Received new tasks:") //Variadic
	variadicTaskLists("Session-1 task", "Session-2 task", "Session-3 task")
}
