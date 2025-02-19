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
		if i%10 == 7 || i%10 == -7 {
			mult *= i
		}
	}
	fmt.Println(mult)
}
