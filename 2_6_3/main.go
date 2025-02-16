package main

import "fmt"

func main() {
	var f float64
	_, _ = fmt.Scan(&f)
	fmt.Println((f - 32) * 5 / 9)
}
