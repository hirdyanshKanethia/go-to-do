package task

import "fmt"

type Task struct {
	ID          int
	Title       string
	Description string
	Done        bool
}

func (t Task) Print() {
	status := "❌"
	if t.Done {
		status = "✅"
	}
	fmt.Printf("[%s] %d: %s - %s\n", status, t.ID, t.Title, t.Description)
}
