package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	num := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&num[i])
	}
	for _, number1 := range num {
		cnt := 0
		for _, number2 := range num {
			if number1 == number2 {
				cnt++
			}
		}
		if cnt == 1 {
			fmt.Print(number1, " ")
		}
	}
}
