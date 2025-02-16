package main

import "fmt"

func main() {
	var x int
	_, _ = fmt.Scan(&x)
	if x%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
