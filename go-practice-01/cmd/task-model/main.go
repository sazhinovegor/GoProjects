package main

import "fmt"

type Task struct {
	ID       int
	Title    string
	Status   string
	Priority int
}



func (t Task) Summary() string {
	// TODO: верните "#ID Название [статус], приоритет N".
	return fmt.Sprintf("#%d %s [%s], приоритет %d", t.ID, t.Title, t.Status, t.Priority)
}

func (t *Task) Complete() {
	t.Status = "done"
}

func (t *Task) SetPriority(priority int) bool {
	if priority < 1 || priority > 3 {
		return false
	}
	t.Priority = priority
	return true
}

func findTaskByID(tasks []Task, id int) *Task {
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i]
		}
	}
	return nil
}

func main() {
	tasks := []Task{
		{ID: 1, Title: "Установить Go", Status: "done", Priority: 1},
		{ID: 2, Title: "Освоить функции", Status: "in_progress", Priority: 2},
		{ID: 3, Title: "Изучить структуры", Status: "new", Priority: 3},
	}
		for i := range tasks {
		if tasks[i].ID == 2 {
			tasks[i].Complete()
		}
	}

	if found := findTaskByID(tasks, 1); found != nil {
		found.Complete()
		found.SetPriority(3)
	}
	completed := 0
	
	for _, task := range tasks {
		fmt.Println(task.Summary())
		if task.Status == "done" {
			completed++
		}
	}
	fmt.Printf("Задач завершено: %d\n", completed)


	if findTaskByID(tasks, 99) == nil {
		fmt.Println("Задача не найдена")
	}

	found := findTaskByID(tasks,2)
	if found != nil {
		found.Complete()
	} 


}