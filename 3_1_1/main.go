package main

import "fmt"

func main() {
	var a int
	var b int
	_, _ = fmt.Scan(&a, &b)
	if a >= b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}
}
