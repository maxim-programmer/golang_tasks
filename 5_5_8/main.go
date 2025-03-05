package main

import "fmt"

func main() {
	var n int
	var h int
	_, _ = fmt.Scan(&n, &h)
	num := make([]int, n)
	for i := range n {
		_, _ = fmt.Scan(&num[i])
	}
	w := 0
	for j := range n {
		w++
		if num[j] > h {
			w++
		}
	}
	fmt.Println(w)
}
