package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Id int
	Title string
	Completed bool
}

func main() {
	printWelcome()

	var taskItems = []Task{
		{Id: 1, Title: "Task 1", Description: "Description 1", Completed: false},
		{Id: 2, Title: "Task 2", Description: "Description 2", Completed: false},
		{Id: 3, Title: "Task 3", Description: "Description 3", Completed: false},
	}

	if len(taskItems) == 0 {
		fmt.Println(" *** No tasks found ***")
	} else {
		printTask(taskItems)
	}

	taskItems = addTask(taskItems)
	printTask(taskItems)
}

func printWelcome() {
	fmt.Println("##########################")
	fmt.Println("## Welcome to task app ##")
	fmt.Println("##########################")
}

func printTask(taskItems []Task) {
	fmt.Println("\nList of tasks:")

	for index, task := range taskItems {
		fmt.Printf("- %d: %s", index+1, task.Title)
	}
}

func addTask(taskItems []Task) []Task {
	fmt.Println("\nadding new task")

	var title string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter task title:")
	title, err := reader.ReadString('\n')
	title = strings.TrimSpace(title)
	if err != nil {
		fmt.Printf("Error reading title: %v\n", err)
		return taskItems
	}
	if title = strings.TrimSpace(title); title == "" {
		fmt.Println("Title cannot be empty")
		return taskItems
	}
	
	taskItems = append(taskItems, Task{
		Id:          len(taskItems) + 1,
		Title:       title,
		Completed:   false,
	})

	return taskItems
}
