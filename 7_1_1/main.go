package main

import "fmt"

func main() {
	var n, k int
	_, _ = fmt.Scan(&n, &k)
	fact_n := factorial(n)
	fact_k := factorial(k)
	fact_nk := factorial(n - k)
	fmt.Println(fact_n / (fact_k * fact_nk))
}

func factorial(number int) int {
	fact := 1
	for i := 1; i <= number; i++ {
		fact *= i
	}
	return fact
}
