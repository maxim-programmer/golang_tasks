package main

import "fmt"

func main() {
	var sum int
	for i := 6; i <= 1000; i++ {
		sum = 0
		for j := 1; j < i; j++ {
			if i%j == 0 {
				sum += j
			}
		}
		if sum == i {
			fmt.Print(i, " ")
		}
	}
}
