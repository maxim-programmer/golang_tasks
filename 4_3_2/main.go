package main

import "fmt"

func main() {
	var (
		a    int
		b    int
		mult = 1
	)
	_, _ = fmt.Scan(&a, &b)
	for i := a; i <= b; i++ {
		mult *= i
	}
	fmt.Println(mult)
}
