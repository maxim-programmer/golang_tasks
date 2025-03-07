package main

import "fmt"

func main() {
	var n, m int
	_, _ = fmt.Scan(&n, &m)
	fmt.Println(umn_razr(n) * umn_razr(m))
}

func umn_razr(x int) int {
	cnt := 0
	for x > 0 {
		cnt++
		x /= 10
	}
	return cnt
}
