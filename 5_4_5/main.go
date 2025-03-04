package main

import "fmt"

func main() {
	var n int
	var m int
	_, _ = fmt.Scan(&n, &m)
	table := make([][]int, n)
	for i := 0; i < n; i++ {
		table[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&table[i][j])
		}
	}
	sum_last := 0
	number := 0
	for x := 0; x < n; x++ {
		sum := 0
		for y := 0; y < m; y++ {
			sum += table[x][y]
		}
		if sum > sum_last {
			sum_last = sum
			number = x
		}
	}
	fmt.Println(sum_last)
	fmt.Println(number)
}
