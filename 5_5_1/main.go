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
	for j := 0; j < n; j++ {
		if num[j]%2 == 0 {
			fmt.Print(num[j], " ")
			cnt++
		}
	}
	fmt.Println()
	fmt.Println(cnt)
}
