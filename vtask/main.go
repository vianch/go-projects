package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Id        int
	Title     string
	Completed bool
}

func main() {
	var taskItems = []Task{
		{Id: 1, Title: "Task 1", Completed: false},
		{Id: 2, Title: "Task 2", Completed: false},
		{Id: 3, Title: "Task 3", Completed: false},
	}

	// Check command line arguments
	args := os.Args[1:]
	if len(args) > 0 {
		switch args[0] {
		case "add":
			taskItems = addTaskFromArgs(taskItems, args)
			printTask(taskItems) // todo: remove this and print just the new task
			return
		case "list":
			printTask(taskItems)
			return	
		case "complete":
			taskItems = completeTaskFromArgs(taskItems, args)
			printTask(taskItems)
			return
		default:
			fmt.Printf("Unknown command: %s\n", args[0])
			return
		}
	} else {
		printWelcome()
	}

	if len(taskItems) == 0 {
		fmt.Println(" *** No tasks found ***")
	} else {
		printTask(taskItems)
	}
}

func printWelcome() {
	fmt.Println("##########################")
	fmt.Println("## Welcome to task app ##")
	fmt.Println("##########################")
}

func printTask(taskItems []Task) {
	fmt.Println("\nList of tasks:")
	
	// Print header
	fmt.Printf("%-5s %-20s %-10s\n", "ID", "Task", "Done")
	
	// Print separator
	fmt.Println(strings.Repeat("-", 35))

	// Print tasks
	for _, task := range taskItems {
		fmt.Printf("%-5d %-20s %-10t\n", task.Id, task.Title, task.Completed)
	}
}

func addTaskFromArgs(taskItems []Task, args []string) []Task {
	if len(args) < 2 {
		fmt.Println("Error: Please provide a task title")
		return taskItems;
	}

	title := args[1]

	if title = strings.TrimSpace(title); title == "" {
		fmt.Println("Title cannot be empty")
		return taskItems
	}

	taskItems = append(taskItems, Task{
		Id:        len(taskItems) + 1,
		Title:     title,
		Completed: false,
	})

	return taskItems
}

func completeTaskFromArgs(taskItems []Task, args []string) []Task {
	if len(args) < 2 {
		fmt.Println("Error: Please provide a task id")
		return taskItems;
	}

	id, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Error: Invalid task id")
		return taskItems;
	}

	if id < 1 || id > len(taskItems) {
		fmt.Println("Error: Task id out of range")
		return taskItems;
	}

	taskItems[id-1].Completed = true
	return taskItems
}