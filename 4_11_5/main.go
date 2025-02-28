package main

import "fmt"

func main() {
	var n int
	var x int
	var min int
	var cnt int = 1
	_, _ = fmt.Scan(&n, &min)
	n--
	for n != 0 {
		_, _ = fmt.Scan(&x)
		if x == min {
			cnt++
		}
		if x < min {
			min = x
			cnt = 1
		}
		n--
	}
	fmt.Println(cnt)
}
