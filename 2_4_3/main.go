package main

import "fmt"

func main() {
	var x int
	_, _ = fmt.Scan(&x)
	x = x * x * x
	fmt.Println(x * x)
}
