package main

import "fmt"

func main() {
	var x int
	_, _ = fmt.Scan(&x)
	fmt.Println((x % 10) + (x / 10 % 10) + (x / 100 % 10))
}
