package main

import "fmt"

func main() {
	var c1 int
	var c2 int
	var c3 int
	_, _ = fmt.Scan(&c1, &c2, &c3)
	fmt.Println(c1/2 + c1%2 + c2/2 + c2%2 + c3/2 + c3%2)
}
