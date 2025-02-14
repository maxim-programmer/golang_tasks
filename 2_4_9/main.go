package main

import "fmt"

func main() {
	var a int
	var b int
	var n int
	_, _ = fmt.Scan(&a, &b, &n)
	fmt.Println((a*n)+((b*n)/100), (b*n)%100)
}
