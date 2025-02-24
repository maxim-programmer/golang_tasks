package main

import "fmt"

func main() {
	var (
		n      int
		max    int
		number int
	)
	_, _ = fmt.Scan(&n, &max)
	for i := 1; i < n; i++ {
		_, _ = fmt.Scan(&number)
		if number > max {
			max = number
		}
	}
	fmt.Println(max)
}
