package main

import "fmt"

func main() {
	var n int
	sum := 0
	_, _ = fmt.Scan(&n)
	for n != 0 {
		if n%2 == 0 && n%3 != 0 {
			sum += n
		}
		_, _ = fmt.Scan(&n)
	}
	fmt.Println(sum)
}
