package main

import "fmt"

func main() {
	var (
		a int
		b int
	)
	_, _ = fmt.Scan(&a, &b)
	for i := b; i >= a; i-- {
		fmt.Println(i)
	}
}
