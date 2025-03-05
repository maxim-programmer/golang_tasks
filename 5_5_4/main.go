package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	num := make([]int, n)
	for i := range n {
		_, _ = fmt.Scan(&num[i])
	}
	cnt := 1
	for j := 1; j < n; j++ {
		if num[j] > num[j-1] {
			cnt++
		}
	}
	fmt.Println(cnt)
}
