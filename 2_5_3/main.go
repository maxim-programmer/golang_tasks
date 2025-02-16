package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	fmt.Println((n % 10) + (n / 10 % 10) + (n / 100))
}
