package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b string
	_, _ = fmt.Scan(&a, &b)
	fmt.Println(obrat(a) + obrat(b))
}

func obrat(x string) int {
	reverse := ""
	for i := len(x) - 1; i >= 0; i-- {
		reverse = reverse + string(x[i])
	}
	x_rev, _ := strconv.Atoi(reverse)
	return x_rev
}
