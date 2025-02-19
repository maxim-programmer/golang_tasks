package main

import "fmt"

func main() {
	var (
		n    int
		mult = 1
	)
	_, _ = fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			mult *= i
		}
	}
	fmt.Println(mult)
}
