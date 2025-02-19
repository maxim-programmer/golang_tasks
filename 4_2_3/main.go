package main

import "fmt"

func main() {
	var (
		n   int
		x   int
		sum = 0
	)
	_, _ = fmt.Scan(&n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&x)
		if x%2 == 0 && x%3 != 0 {
			sum += x
		}
	}
	fmt.Println(sum)
}
