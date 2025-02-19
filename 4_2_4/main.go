package main

import "fmt"

func main() {
	var (
		n     int
		x     int
		count int
	)
	_, _ = fmt.Scan(&n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&x)
		if x%10 == 0 {
			count++
		}
	}
	fmt.Println(count)
}
