package main

import "fmt"

func main() {
	var a, b int
	_, _ = fmt.Scan(&a, &b)
	fmt.Println(sign(a) + sign(b))
}

func sign(x int) int {
	if x < 0 {
		return -1
	} else if x == 0 {
		return 0
	} else {
		return 1
	}
}
