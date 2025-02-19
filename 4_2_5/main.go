package main

import "fmt"

func main() {
	var (
		n    int
		x    int
		flag = "NO"
	)
	_, _ = fmt.Scan(&n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&x)
		if x == 0 {
			flag = "YES"
		}
	}
	fmt.Println(flag)
}
