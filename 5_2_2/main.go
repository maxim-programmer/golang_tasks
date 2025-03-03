package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	num := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&num[i])
	}
	min := num[0]
	for j := 1; j < n; j++ {
		if num[j] < min {
			min = num[j]
		}
	}
	for k := 0; k < n; k++ {
		num[k] -= min
		fmt.Print(num[k], " ")
	}
}
