package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	for i := 2; i < n; i++ {
		if n%i == 0 {
			fmt.Println(n / i)
			break
		}
	}
}
