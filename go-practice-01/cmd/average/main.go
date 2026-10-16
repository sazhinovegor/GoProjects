package main

import "fmt"

func main() {
	var numFirst int
	var numSecond int
	var numThird int
	fmt.Scan(&numFirst)
	fmt.Scan(&numSecond)
	fmt.Scan(&numThird)
	var sum int = numFirst + numSecond + numThird
	var average float64 = float64(sum) / 3.0
	fmt.Println("Сумма:", sum)
	fmt.Printf("Среднее значение: %.2f\n", average)
}
