package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	num := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&num[i])
	}
	for j := 0; j < n; j++ {
		if j+1 < n && j%2 == 0 {
			num[j], num[j+1] = num[j+1], num[j]
		}
	}
	for k := 0; k < n; k++ {
		fmt.Print(num[k], " ")
	}
}
