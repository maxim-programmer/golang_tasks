package main

import "fmt"

func main() {
	var (
		max1 int
		max2 int
		n    int
	)
	_, _ = fmt.Scan(&n, max2)
	max1 = n
	if max2 > max1 {
		max1, max2 = max2, max1
	}
	for n != 0 {
		_, _ = fmt.Scan(&n)
		if n > max1 {
			max2 = max1
			max1 = n
		} else if n > max2 {
			max2 = n
		}
	}
	fmt.Println(max2)
}
