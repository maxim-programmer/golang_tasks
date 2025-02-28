package main

import "fmt"

func main() {
	var (
		a int
		n int
		x int = 1
	)
	_, _ = fmt.Scan(&a, &n)
	for n > 0 {
		x *= a
		n--
	}
	fmt.Println(x)
}
