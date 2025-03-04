package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			_, _ = fmt.Scan(&matrix[i][j])
		}
	}
	flag := true
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if !(x == y || matrix[x][y] == matrix[y][x]) {
				flag = false
				break
			}
		}
	}
	if flag {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
