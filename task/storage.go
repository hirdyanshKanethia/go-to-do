package task

import (
	"encoding/json"
	"os"
	"fmt"
)

func SaveTasks(filename string, tasks []Task) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(tasks)
}

func LoadTasks(filename string) ([]Task, error) {
	var tasks []Task
	
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return tasks, nil
		}
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	return tasks, err
}

func CreateTask(filename string, task Task) error {
	tasks, err := LoadTasks(filename)
	if err != nil {
		return err
	}

	for _, t := range tasks {
		if t.Title == task.Title {
			return fmt.Errorf("[Error] A task with the same name already exists!")
		}
	}

	tasks = append(tasks, task)
	return SaveTasks(filename, tasks)
}
