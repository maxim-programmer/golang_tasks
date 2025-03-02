package main

import "fmt"

func main() {
	var n int
	var cnt int
	_, _ = fmt.Scan(&n)
	num := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&num[i])
	}
	for j := 1; j < n; j++ {
		if num[j] > num[j-1] {
			cnt++
		}
	}
	fmt.Println(cnt)
}
