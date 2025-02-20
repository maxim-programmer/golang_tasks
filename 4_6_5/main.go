package main

import "fmt"

func main() {
	var n int
	cnt := 0
	_, _ = fmt.Scan(&n)
	last := 0
	for n != 0 {
		if last*n < 0 {
			cnt++
		}
		last = n
		_, _ = fmt.Scan(&n)
	}
	fmt.Println(cnt)
}
