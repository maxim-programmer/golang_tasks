package main

import "fmt"

func main() {
	var radius float64
	_, _ = fmt.Scan(&radius)
	fmt.Println(3.14 * radius * radius)
}
