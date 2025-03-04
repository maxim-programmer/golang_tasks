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
			table[i][j] = (j + 1) * (i + 1)
		}
	}
	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			fmt.Print(table[x][y], " ")
		}
		fmt.Println()
	}
}
