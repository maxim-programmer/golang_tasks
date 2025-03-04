package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	table := make([][]int, n)
	for i := 0; i < n; i++ {
		table[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == n-1-j {
				table[i][j] = 1
			} else if i > n-1-j {
				table[i][j] = 2
			} else {
				table[i][j] = 0
			}
		}
	}
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			fmt.Print(table[x][y], " ")
		}
		fmt.Println()
	}
}
