package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	last_digit := n % 10
	if last_digit%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
