package main

import "fmt"

func main() {
	var a, b, c int
	_, _ = fmt.Scan(&a, &b, &c)
	fmt.Println(fact2(a), fact2(b), fact2(c))
}

func fact2(n int) int {
	mult := 1
	if n%2 == 1 {
		for i := 1; i <= n; i += 2 {
			mult *= i
		}
	} else {
		for i := 2; i <= n; i += 2 {
			mult *= i
		}
	}
	return mult
}
