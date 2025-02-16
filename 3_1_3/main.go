package main

import "fmt"

func main() {
	var age int
	_, _ = fmt.Scan(&age)
	if age >= 12 {
		fmt.Println("Доступ разрешен")
	} else {
		fmt.Println("Доступ запрещен")
	}
}
