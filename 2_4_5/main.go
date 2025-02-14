package main

import "fmt"

func main() {
	var student_count int
	var apple_count int
	_, _ = fmt.Scan(&student_count, &apple_count)
	fmt.Println(apple_count / student_count)
}
