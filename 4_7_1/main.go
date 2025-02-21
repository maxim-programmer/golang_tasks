package main

import "fmt"

func main() {
	var (
		a   int
		b   int
		k   int
		cnt int
	)
	_, _ = fmt.Scan(&a, &b, &k)
	for i := a; i <= b; i++ {
		cnt = 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				cnt++
			}
		}
		if cnt >= k {
			fmt.Print(i)
			fmt.Print(" ")
		}

	}
}
