package main

import "fmt"

func main() {
	var d1 int
	var d2 int
	var d3 int
	_, _ = fmt.Scan(&d1, &d2, &d3)
	p1 := d1*2 + d2*2
	p2 := d1 + d3 + d2
	p3 := d1*2 + d3*2
	p4 := d2*2 + d3*2
	fmt.Println(min(p1, p2, p3, p4))
}
