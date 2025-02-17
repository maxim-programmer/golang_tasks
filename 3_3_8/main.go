package main

import "fmt"

func main() {
	var month int
	_, _ = fmt.Scan(&month)
	switch month {
	case 12, 1, 2:
		fmt.Println("Зима")
	case 3, 4, 5:
		fmt.Println("Весна")
	case 6, 7, 8:
		fmt.Println("Лето")
	default:
		fmt.Println("Осень")
	}
}
