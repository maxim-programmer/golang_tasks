package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	for n <= 100 {
		if n >= 10 {
			fmt.Println(n)
		}
		_, _ = fmt.Scan(&n)
	}
}
