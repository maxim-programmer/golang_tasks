package main

import "fmt"

func main() {
	var a, b int
	_, _ = fmt.Scan(&a, &b)
	sum_a := sum_digits(a)
	sum_b := sum_digits(b)
	if sum_a > sum_b {
		fmt.Println(1)
	} else if sum_a < sum_b {
		fmt.Println(2)
	} else {
		fmt.Println(0)
	}
}

func sum_digits(x int) int {
	sum := 0
	for x > 0 {
		sum += x % 10
		x /= 10
	}
	return sum
}
