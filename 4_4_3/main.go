package main

import "fmt"

func main() {
	var n int
	var pokazatel = 0
	var stepen = 1
	_, _ = fmt.Scan(&n)
	for stepen < n {
		pokazatel++
		stepen *= 2
	}
	fmt.Println(pokazatel)
}
