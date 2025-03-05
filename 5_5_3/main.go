package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	num := make([]int, n)
	for i := range n {
		_, _ = fmt.Scan(&num[i])
	}
	cnt := 0
	for j := range n {
		if num[j]%3 == 0 && num[j]%10 == 7 {
			cnt++
		}
	}
	for k := range n {
		if num[k]%3 == 0 && num[k]%10 == 7 {
			num[k] = cnt
		}
		fmt.Print(num[k], " ")
	}
}
