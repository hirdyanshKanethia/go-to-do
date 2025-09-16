package task

import "fmt"

type Task struct {
	ID    int
	Title string
	Done  bool
}

func (t Task) Print() {
	status := "X"
	if t.Done {
		status = "*"
	}
	fmt.Printf("[%s] %d: %s\n", status, t.ID, t.Title)
}
