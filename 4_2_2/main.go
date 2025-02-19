package main

import "fmt"

func main() {
	var (
		n   int
		sum = 0
		x   int
	)
	_, _ = fmt.Scan(&n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&x)
		sum += x
	}
	fmt.Println(sum)
}
