package main

import "fmt"

func main() {
	var a float64
	var b float64
	_, _ = fmt.Scan(&a, &b)
	fmt.Println((a + b) / 2)
}
