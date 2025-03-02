package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	num := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&num[i])
	}
	x := num[0]
	for j := 0; j < n; j++ {
		x, num[(j+1)%n] = num[(j+1)%n], x
	}
	for k := 0; k < n; k++ {
		fmt.Print(num[k], " ")
	}
}
