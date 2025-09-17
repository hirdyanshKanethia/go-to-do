package task

import (
	"fmt"
)

type HistoryManager struct {
	undoStack map[int][]string
	redoStack map[int][]string
}

func NewHistoryManager() *HistoryManager {
	return &HistoryManager{
		undoStack: make(map[int][]string),
		redoStack: make(map[int][]string),
	}
}

func UndoTask(ID int, tasks []Task, hm *HistoryManager) error {
	for i := range tasks {
		if tasks[i].ID == ID {
			if len(hm.undoStack[ID]) == 0 {
				fmt.Printf("Already at first change for task ID: %d\n", ID)
				return fmt.Errorf("nothing to undo for task ID: %d", ID)
			}
			hm.redoStack[ID] = append(hm.redoStack[ID], tasks[i].Description)

			lastIndex := len(hm.undoStack[ID]) - 1
			tasks[i].Description = hm.undoStack[ID][lastIndex]
			hm.undoStack[ID] = hm.undoStack[ID][:lastIndex]

			return nil
		}
	}
	return fmt.Errorf("[ERROR] Could not find task with the given ID: %d", ID)
}

func RedoTask(ID int, tasks []Task, hm *HistoryManager) error {
	for i := range tasks {
		if tasks[i].ID == ID {
			if len(hm.redoStack[ID]) == 0 {
				fmt.Printf("Already at latest change for task ID: %d\n", ID)
				return fmt.Errorf("nothing to redo for task ID: %d", ID)
			}
			hm.undoStack[ID] = append(hm.undoStack[ID], tasks[i].Description)

			lastIndex := len(hm.redoStack[ID]) - 1
			tasks[i].Description = hm.redoStack[ID][lastIndex]
			hm.redoStack[ID] = hm.redoStack[ID][:lastIndex]

			return nil
		}
	}
	return fmt.Errorf("[ERROR] Could not find task with the given ID: %d", ID)
}

func ClearRedoStack(ID int, hm *HistoryManager) {
	hm.redoStack[ID] = hm.redoStack[ID][:0]
}

func EditTask(ID int, tasks []Task, newDesc string, hm *HistoryManager) {
	for i := range tasks {
		if tasks[i].ID == ID {
			hm.undoStack[ID] = append(hm.undoStack[ID], tasks[i].Description)
			tasks[i].Description = newDesc
			hm.redoStack[ID] = hm.redoStack[ID][:0]
		}
	}
}
