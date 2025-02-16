package main

import "fmt"

func main() {
	var x int
	_, _ = fmt.Scan(&x)
	fmt.Println(x / 100 % 10)
}
