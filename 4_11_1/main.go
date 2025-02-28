package main

import "fmt"

func main() {
	var n int
	var x int
	var cnt int
	_, _ = fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&x)
		if x == 0 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
