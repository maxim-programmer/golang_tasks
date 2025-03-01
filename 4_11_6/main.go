package main

import "fmt"

func main() {
	var (
		n   int
		ck  int = 0
		ck1 int = 0
		ck2 int = 0
	)
	_, _ = fmt.Scan(&n)
	for n > 0 {
		ck += n % 10
		n /= 10
	}
	for ck > 0 {
		ck1 += ck % 10
		ck /= 10
	}
	if ck1 >= 10 {
		for ck1 > 0 {
			ck2 += ck1 % 10
			ck1 /= 10
		}
	}
	if ck1 < 10 && ck2 == 0 {
		fmt.Println(ck1)
	} else {
		fmt.Println(ck2)
	}
}
