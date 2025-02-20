package main

import "fmt"

func main() {
	var (
		n      int
		stepen = 1
	)
	_, _ = fmt.Scan(&n)
	for stepen <= n {
		fmt.Println(stepen)
		stepen *= 2
	}
}
