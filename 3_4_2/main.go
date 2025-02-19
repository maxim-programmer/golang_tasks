package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	if n%2 == 0 {
		fmt.Println(n + 2)
	} else {
		fmt.Println(n + 1)
	}
}
