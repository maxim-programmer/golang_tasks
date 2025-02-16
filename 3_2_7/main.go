package main

import "fmt"

func main() {
	var age int
	_, _ = fmt.Scan(&age)
	if age <= 13 {
		fmt.Println("детство")
	} else if age >= 14 && age <= 24 {
		fmt.Println("молодость")
	} else if age >= 25 && age <= 59 {
		fmt.Println("зрелость")
	} else {
		fmt.Println("старость")
	}
}
