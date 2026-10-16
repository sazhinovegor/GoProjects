package main

import "fmt"

func main() {
	var minutes int
	fmt.Scan(&minutes)
	var hour int = minutes / 60
	var ostminutes int = minutes - 60*hour
	if minutes != nil {
		fmt.Println("Ошибка! Введите, пожалуйста, целое число.")
		return

	}
	if minutes < 0 {
		fmt.Println("Ошибка! Введите, пожалуйста, положительное число.")
	}
	if minutes.(string) {
		fmt.Println("Ошибка! Введите, пожалуйста, числовое значение.")
	}

	fmt.Println("Полных часов", hour)
	fmt.Println("Оставшихся минут", ostminutes)
}
