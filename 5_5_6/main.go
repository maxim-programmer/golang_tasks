package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	num := make([]int, n)
	for i := range n {
		_, _ = fmt.Scan(&num[i])
	}
	for j := range n {
		for k := range n {
			if num[j]+num[k] == 0 && j < k {
				fmt.Println(j, k)
			}
		}
	}
}
