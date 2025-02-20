package main

import "fmt"

func main() {
	var n int
	sum := 0
	_, _ = fmt.Scan(&n)
	x := n
	for x != 0 {
		sum += x % 10
		x /= 10
	}
	if n%sum == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
