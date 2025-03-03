package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	num := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&num[i])
	}
	MaxIndex := 0
	MinIndex := 0
	for j := 1; j < n; j++ {
		if num[j] < num[MinIndex] {
			MinIndex = j
		}
		if num[j] > num[MaxIndex] {
			MaxIndex = j
		}
	}
	fmt.Println(MaxIndex - MinIndex)
}
