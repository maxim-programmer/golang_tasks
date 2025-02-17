package main

import "fmt"

func main() {
	var weight int
	_, _ = fmt.Scan(&weight)
	if weight < 60 {
		fmt.Println("Легкий вес")
	} else {
		if weight < 64 {
			fmt.Println("Первый полусредний вес")
		} else {
			fmt.Println("Полусредний вес")
		}
	}
}
