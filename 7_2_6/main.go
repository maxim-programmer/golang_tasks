package main

import "fmt"

func main() {
	var a, b int
	_, _ = fmt.Scan(&a, &b)
	cnt1 := cnt_digits(a)
	cnt2 := cnt_digits(b)
	if cnt1 > cnt2 {
		fmt.Println(1)
	} else if cnt1 < cnt2 {
		fmt.Println(2)
	} else {
		fmt.Println(0)
	}
}

func cnt_digits(x int) int {
	cnt := 0
	for x > 0 {
		cnt++
		x /= 10
	}
	return cnt
}
