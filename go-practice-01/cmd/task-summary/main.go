package main

import "fmt"

func statusLabel(status string) string {
	switch status {
	case "new":
		return "Новая"
	case "in_progress":
		return "В работе"
	case "done":
		return "Завершена"
	default:
		return "Неизвестный статус"
	}
}

// countStatus считает количество элементов с заданным статусом.
func countStatus(statuses []string, target string) int {
	count := 0
	if len(statuses) == 0 {
		return 0
	} else {
		for i := 0; i < len(statuses); i++ {
			if statuses[i] == target {
				count++
			}
		}
		return count
	}

}

func main() {
	statuses := []string{"new", "in_progress", "done", "done"}
	//statuses = append(statuses,"new")

	fmt.Printf("Задач всего: %d\n", len(statuses))
	for i, status := range statuses {
		fmt.Printf("%d. %s\n", i+1, statusLabel(status))
	}
	fmt.Printf("Завершено: %d\n", countStatus(statuses, "done"))

	counts := make(map[string]int)
	for _, status := range statuses {
		counts[status]++
	}

	fmt.Printf("Статистика: new=%d, in_progress=%d, done=%d", counts["new"], counts["in_progress"], counts["done"])
}
