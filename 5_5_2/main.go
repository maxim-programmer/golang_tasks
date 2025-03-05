package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	num := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&num[i])
	}
	cnt := 0
	for j := 0; j < n-1; j++ {
		if num[j] < num[j+1] {
			cnt++
		}
	}
	fmt.Println(cnt)
}
