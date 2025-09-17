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

	// HistoryManager Initialization
	hm := task.NewHistoryManager()

	for {
		fmt.Println("\nCommands: add/edit/undo/redo/list/done/delete/exit")
		fmt.Print("Enter command: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch {
		case strings.HasPrefix(input, "add"):
			info := strings.SplitN(strings.TrimSpace(strings.TrimPrefix(input, "add")), "-", 2)
			if len(info) < 2 {
				fmt.Printf("[ERROR] Invalid add command. Please use the add command with the proper syntax\n")
				break
			}
			title, desc := info[0], info[1]
			title = strings.TrimSpace(title)
			desc = strings.TrimSpace(desc)
			newTask := task.Task{
				ID:          len(tasks) + 1,
				Title:       title,
				Description: desc,
				Done:        false,
			}

			err := task.CreateTask(filename, newTask)
			if err != nil {
				fmt.Println(err)
			}

			tasks = append(tasks, newTask)
			fmt.Println("Task added!")

		case strings.HasPrefix(input, "edit"):
			info := strings.SplitN(strings.TrimSpace(strings.TrimPrefix(input, "edit")), "-", 2)
			if len(info) != 2 {
				fmt.Printf("[ERROR] Invalid add command. Please use the add command with the proper syntax\n")
				break
			}
			id, err := strconv.Atoi(strings.TrimSpace(info[0]))
			if err != nil {
				fmt.Printf("[ERROR] %v\n", err)
				break
			}
			newDesc := strings.TrimSpace(info[1])
			task.EditTask(id, tasks, newDesc, hm)
			task.SaveTasks(filename, tasks)
			fmt.Printf("edit successful for task ID: %d\n", id)

		case strings.HasPrefix(input, "undo"):
			info := strings.TrimPrefix(input, "undo")
			id, err := strconv.Atoi(strings.TrimSpace(info))
			if err != nil {
				fmt.Printf("[ERROR] %v\n", err)
				break
			}
			task.UndoTask(id, tasks, hm)
			fmt.Printf("Undo successful for task ID: %d\n", id)

		case strings.HasPrefix(input, "redo"):
			info := strings.TrimPrefix(input, "redo")
			id, err := strconv.Atoi(strings.TrimSpace(info))
			if err != nil {
				fmt.Printf("[ERROR] %v\n", err)
			}
			task.RedoTask(id, tasks, hm)
			fmt.Printf("Redo successful for task ID: %d\n", id)

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
			fmt.Println("[ERROR] Invalid Command")
		}
	}
}
