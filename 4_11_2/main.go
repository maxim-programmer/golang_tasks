package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	var s string
	_, _ = fmt.Scan(&n)
	for i := n; i > 0; {
		s += strconv.Itoa(i % 10)
		i /= 10
	}
	if strconv.Itoa(n) == s {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
