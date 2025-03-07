package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	var s string
	u := make([]string, n)
	for i := range n {
		_, _ = fmt.Scan(&s)
		if len(s) > 10 {
			u[i] = string(s[0]) + strconv.Itoa(len(s)-2) + string(s[len(s)-1])
		} else {
			u[i] = s
		}
	}
	for j := range n {
		fmt.Println(u[j])
	}
}
