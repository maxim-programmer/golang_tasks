package main

import "fmt"

func main() {
	var x int
	_, _ = fmt.Scan(&x)
	if x >= -3 && x <= 1 {
		fmt.Println("YES")
	} else {
		if x >= 5 && x <= 9 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
