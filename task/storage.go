package task

import (
	"encoding/json"
	"os"
)

func SaveTasks(filename string, tasks []Task) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
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
