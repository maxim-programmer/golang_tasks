package main

import "fmt"

func main() {
	var n int
	var s string
	_, _ = fmt.Scan(&n, &s)
	cnt := 0
	count := 0
	for i := range n {
		if string(s[i]) == "x" {
			cnt++
		}
		if (cnt >= 3 && string(s[i]) != "x") || (cnt >= 3 && i == len(s)-1) {
			count += cnt - 2
			cnt = 0
		}
		if string(s[i]) != "x" {
			cnt = 0
		}
	}
	fmt.Println(count)
}
