package main

import "fmt"

func main() {
	var size int
	_, _ = fmt.Scan(&size)
	num := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&num[i])
	}
	for j := 0; j < size; j++ {
		if num[j]%3 == 0 {
			fmt.Print(num[j], " ")
		}
	}
}
