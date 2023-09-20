package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	tasks := [] string {}

	for {
		fmt.Println("To-Do List: ")
		fmt.Print("\n Options:\n 1. Add Task \n 2. Remove Task \n 3. Quit \n Enter your choice: ")
		
		var choice int
		fmt.Scanln(&choice)

		switch choice {
			case 1:
				addTask(&tasks)
			case 2:
				removeTask(&tasks)
			case 3:
				fmt.Println("Good bye")
				os.Exit(0)
			default:
				fmt.Println("Invalid choice. Please try again")
		}
	}
}

func addTask(tasks *[]string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(" Enter task: ")
	scanner.Scan()
	task := scanner.Text()

	if task != "" {
		*tasks = append(*tasks, task)
		fmt.Printf("Task \"%s\" added.\n", task)
	}
}


func listTask(tasks []string) {
	if len(tasks) == 0 {
		fmt.Println("No Tasks.")
		return
	}
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func removeTask(tasks *[]string) {
	listTask(*tasks)
	fmt.Print("Enter the number of the task to remove: ")

	var num int
	fmt.Scanln(&num)

	if num >= 1 && num <= len(*tasks) {
		removeTask := (*tasks)[num - 1]
		*tasks = append((*tasks)[:num - 1], (*tasks)[num:]...)
		fmt.Printf("Task \"%s\" removed.\n", removeTask)
	} else {
		fmt.Println("Invalid task number.")
	}
}

