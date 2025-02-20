package main

import "fmt"

func main() {
	var a int
	var b int
	_, _ = fmt.Scan(&a, &b)
	for a != 0 && b != 0 {
		if a > b {
			a = a % b
		} else {
			b = b % a
		}

	}
	fmt.Println(a + b)
}
