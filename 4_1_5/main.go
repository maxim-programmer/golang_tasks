package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Println(i * i)
	}
}
