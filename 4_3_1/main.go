package main

import "fmt"

func main() {
	var (
		n    int
		fact = 1
	)
	_, _ = fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fact *= i
	}
	fmt.Println(fact)
}
