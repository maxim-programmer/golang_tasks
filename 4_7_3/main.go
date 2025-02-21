package main

import "fmt"

func main() {
	var (
		n        int
		sum_last int = 0
		n_last   int
		sum      int = 0
		num      int
	)
	_, _ = fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		num = i
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		if sum > sum_last {
			sum_last = sum
			n_last = i
		}
		sum = 0
	}
	fmt.Println(n_last, sum_last)
}
