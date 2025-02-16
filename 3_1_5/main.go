package main

import "fmt"

func main() {
	var a int
	var b int
	_, _ = fmt.Scan(&a, &b)
	if (a % b) == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
