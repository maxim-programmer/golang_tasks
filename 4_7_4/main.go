package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	for i := 2; i <= n; i++ {
		for n%i == 0 && (i%2 != 0 || i == 2) {
			fmt.Print(i)
			fmt.Print(" ")
			n /= i
		}
	}
}
