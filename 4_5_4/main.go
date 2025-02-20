package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	for n != 0 {
		fmt.Print(n % 2)
		n /= 2
	}
}
