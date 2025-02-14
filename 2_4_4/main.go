package main

import "fmt"

func main() {
	var a int
	var b int
	var c int
	_, _ = fmt.Scan(&a, &b, &c)
	fmt.Println(a * b * c)
}
