package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	num := make([]int, n)
	for i := range n {
		_, _ = fmt.Scan(&num[i])
	}
	var rost int
	_, _ = fmt.Scan(&rost)
	number := 1
	for j := range n {
		if num[j] >= rost {
			number++
		}
	}
	fmt.Println(number)
}
