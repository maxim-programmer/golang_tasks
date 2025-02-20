package main

import "fmt"

func main() {
	var n int
	cnt := 0
	_, _ = fmt.Scan(&n)
	for n != 0 {
		if n%10 == 4 {
			cnt++
		}
		n /= 10
	}
	fmt.Println(cnt)
}
