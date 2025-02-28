package main

import "fmt"

func main() {
	var n int
	var ten int = 1
	var x int
	_, _ = fmt.Scan(&n)
	for n > 0 {
		if (n%10 != 5) && (n%10 != 7) {
			x += n % 10 * ten
			ten *= 10
		}
		n /= 10
	}
	fmt.Println(x)
}
