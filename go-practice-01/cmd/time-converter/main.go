package main

import "fmt"

func main() {
	var minutes int
	_, err := fmt.Scan(&minutes)
	if minutes < 0 {
		fmt.Println("Ошибка! Введите, пожалуйста, положительное число.")
		return
	}
	if err != nil {
		fmt.Println("Ошибка! Введите, пожалуйста, целое значение.")
		return
	}

	var hour int = minutes / 60
	var ostminutes int = minutes - 60*hour

	fmt.Println("Полных часов", hour)
	fmt.Println("Оставшихся минут", ostminutes)
}
