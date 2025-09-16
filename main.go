package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"todo-app/task"
)

func main() {
	const filename = "tasks.json"
	tasks, _ := task.LoadTasks(filename)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nCommands: add/list/done/delete/exit")
		fmt.Print("Enter command: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch {
		case strings.HasPrefix(input, "add"):
			title := strings.TrimSpace(strings.TrimPrefix(input, "add"))
			newTask := task.Task{
				ID:    len(tasks) + 1,
				Title: title,
				Done:  false,
			}
			tasks = append(tasks, newTask)
			task.SaveTasks(filename, tasks)
			fmt.Println("Task added!")

		case input == "list":
			for _, t := range tasks {
				t.Print()
			}

		case strings.HasPrefix(input, "done"):
			idStr := strings.TrimSpace(strings.TrimPrefix(input, "done"))
			id, _ := strconv.Atoi(idStr)
			for i := range tasks {
				if tasks[i].ID == id {
					tasks[i].Done = true
				}
			}
			task.SaveTasks(filename, tasks)
			fmt.Println("Task marked as done!")

		case strings.HasPrefix(input, "delete"):
			idStr := strings.TrimSpace(strings.TrimPrefix(input, "delete"))
			id, _ := strconv.Atoi(idStr)
			newTasks := []task.Task{}
			for _, t := range tasks {
				if t.ID != id {
					newTasks = append(newTasks, t)
				}
			}
			tasks = newTasks
			task.SaveTasks(filename, tasks)
			fmt.Println("Task deleted!")

		case input == "exit":
			fmt.Println("Bye!")
			return

		default:
			fmt.Println("Invalid Command")
		}
	}
}
