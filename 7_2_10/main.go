package main

import "fmt"

func main() {
	var a, b, c int
	_, _ = fmt.Scan(&a, &b, &c)
	fmt.Println(SumRange(a, b) + SumRange(b, c))
}

func SumRange(x int, y int) int {
	if x > y {
		return 0
	}
	sum := 0
	for i := x; i <= y; i++ {
		sum += i
	}
	return sum
}
