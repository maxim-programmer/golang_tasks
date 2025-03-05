package main

import "fmt"

func main() {
	var n int
	var m int
	_, _ = fmt.Scan(&n, &m)
	table := make([][]string, n)
	for i := range n {
		table[i] = make([]string, m)
		for j := range m {
			_, _ = fmt.Scan(&table[i][j])
		}
	}
	flag := false
	for x := range n {
		for y := range m {
			if table[x][y] == "W" || table[x][y] == "B" || table[x][y] == "G" {
				flag = true
			} else {
				flag = false
				break
			}
		}
		if flag == false {
			break
		}
	}
	if flag {
		fmt.Println("#Black&White")
	} else {
		fmt.Println("#Color")
	}
}
