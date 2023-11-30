package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Task represents a to-do list item
type Task struct {
	ID       int
	Text     string
	Completed bool
}

// TodoList represents a collection of tasks
type TodoList struct {
	Tasks []Task
}

// AddTask adds a new task to the to-do list
func (tl *TodoList) AddTask(text string) {
	newTask := Task{
		ID:       len(tl.Tasks) + 1,
		Text:     text,
		Completed: false,
	}
	tl.Tasks = append(tl.Tasks, newTask)
	fmt.Printf("Task added: %s\n", text)
}

// MarkCompleted marks a task as completed
func (tl *TodoList) MarkCompleted(id int) {
	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks[i].Completed = true
			fmt.Printf("Task marked as completed: %s\n", task.Text)
			return
		}
	}
	fmt.Println("Task not found.")
}

// ViewTasks displays the current to-do list
func (tl *TodoList) ViewTasks() {
	if len(tl.Tasks) == 0 {
		fmt.Println("No tasks in the to-do list.")
		return
	}

	fmt.Println("To-Do List:")
	for _, task := range tl.Tasks {
		status := "Incomplete"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("%d. %s - %s\n", task.ID, task.Text, status)
	}
}

// getUserInput gets user input from the console
func getUserInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func main() {
	todoList := TodoList{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the To-Do List App!")

	for {
		fmt.Println("\nOptions:")
		fmt.Println("1. Add Task")
		fmt.Println("2. Mark Task as Completed")
		fmt.Println("3. View Tasks")
		fmt.Println("4. Exit")

		choice, err := getUserInput("Enter your choice (1-4): ", reader)
		if err != nil {
			fmt.Println("Error getting user input:", err)
			continue
		}

		switch choice {
		case "1":
			taskText, err := getUserInput("Enter the task: ", reader)
			if err != nil {
				fmt.Println("Error getting task input:", err)
				continue
			}
			todoList.AddTask(taskText)

		case "2":
			idStr, err := getUserInput("Enter the task ID to mark as completed: ", reader)
			if err != nil {
				fmt.Println("Error getting ID input:", err)
				continue
			}
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid ID. Please enter a number.")
				continue
			}
			todoList.MarkCompleted(id)

		case "3":
			todoList.ViewTasks()

		case "4":
			fmt.Println("Exiting the To-Do List App. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
		}
	}
}
