package main

import "fmt"

func main() {
	var n int
	cnt := 0
	_, _ = fmt.Scan(&n)
	last := n
	for n != 0 {
		_, _ = fmt.Scan(&n)
		if n > last {
			cnt++
		}
		last = n
	}
	fmt.Println(cnt)
}
