package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	num := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&num[i])
	}
	min_index := 0
	min := num[min_index]
	for j := 1; j < n; j++ {
		if num[j] < min {
			min = num[j]
			min_index = j
		}
	}
	fmt.Println(min_index)
}
