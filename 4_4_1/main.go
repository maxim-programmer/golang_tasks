package main

import "fmt"

func main() {
	var (
		a   int
		cnt = 0
	)
	_, _ = fmt.Scan(&a)
	for a%3 == 0 {
		cnt++
		a = a / 3
	}
	fmt.Println(cnt)
}
