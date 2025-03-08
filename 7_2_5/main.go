package main

import "fmt"

func main() {
	var n, m int
	_, _ = fmt.Scan(&n)
	a := massiv(n)
	_, _ = fmt.Scan(&m)
	b := massiv(m)
	fmt.Println(max(a, n) * max(b, m))
}

func max(x []int, k int) int {
	max_x := x[0]
	for i := 1; i < k; i++ {
		if max_x < x[i] {
			max_x = x[i]
		}
	}
	return max_x
}
func massiv(cnt int) []int {
	m := make([]int, cnt)
	for i := range cnt {
		_, _ = fmt.Scan(&m[i])
	}
	return m
}
