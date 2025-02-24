package main

import "fmt"

func main() {
	var (
		n      int
		max    int
		min    int
		number int
	)
	_, _ = fmt.Scan(&n, &max)
	min = max
	for i := 1; i < n; i++ {
		_, _ = fmt.Scan(&number)
		if number > max {
			max = number
		} else if number < min {
			min = number
		}
	}
	fmt.Print(max, " ")
	if min < 30 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
