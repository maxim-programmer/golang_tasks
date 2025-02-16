package main

import "fmt"

func main() {
	var x int
	_, _ = fmt.Scan(&x)
	if x >= 100 && x <= 999 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
