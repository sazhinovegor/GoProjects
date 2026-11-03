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
	// TODO: верните &tasks[i] при совпадении, иначе nil.
	return nil
}

func main() {
	taskFirst := Task{ID: 1, Title: "Изучить GO", Status: "New", Priority: 2}
	taskSecond := Task{ID:2, Title: "Прийти домой", Status: "Old", Priority: 1}
	taskThird := Task{ID:3, Title: "Вкусно покушать", Status:"New",Priority: 1}
	fmt.Println(taskFirst)
	fmt.Println(taskSecond)
	fmt.Println(taskThird)
	ptr := &taskFirst
	ptr.Priority = 3
	fmt.Println(taskFirst.Priority)
	tasks := []Task{
		{ID: 1, Title: "Установить Go", Status: "done", Priority: 1},
		{ID: 2, Title: "Освоить функции", Status: "in_progress", Priority: 2},
		{ID: 3, Title: "Изучить структуры", Status: "new", Priority: 3},
	}

	// TODO: найдите ID=2, завершите и установите приоритет 3.
	// TODO: посчитайте завершённые задачи.
		for i := range tasks {
		if tasks[i].ID == 2 {
			tasks[i].Complete()
		}
	}
	
	for _, task := range tasks {
		fmt.Println(task.Summary())
	}
	if findTaskByID(tasks, 99) == nil {
		fmt.Println("Задача не найдена")
	}
}