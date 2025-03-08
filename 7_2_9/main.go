package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	fmt.Println(prime(n))
}

func prime(x int) string {
	cnt := 0
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			cnt++
		}
	}
	if cnt == 2 {
		return "prime"
	}
	return "composite"
}
