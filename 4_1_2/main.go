package main

import "fmt"

func main() {
	var (
		a int
		b int
	)
	_, _ = fmt.Scan(&a, &b)
	for i := a; i <= b; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
