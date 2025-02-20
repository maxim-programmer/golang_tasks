package main

import "fmt"

func main() {
	var n int
	cnt := 0
	_, _ = fmt.Scan(&n)
	for n != 0 {
		if n%2 == 1 {
			cnt++
		}
		n /= 2
	}
	fmt.Println(cnt)
}
