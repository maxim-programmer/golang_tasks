package main

import "fmt"

func main() {
	var k2 int
	var k3 int
	var k5 int
	var k6 int
	_, _ = fmt.Scan(&k2, &k3, &k5, &k6)
	var s int
	var c2 int
	if k2 <= k5 && k2 <= k6 {
		s += 256 * k2
		c2 = k2
	} else if k5 <= k2 && k5 <= k6 {
		s += 256 * k5
		c2 = k5
	} else if k6 <= k2 && k6 <= k5 {
		s += 256 * k6
		c2 = k6
	}
	if (k2 - c2) <= k3 {
		s += 32 * (k2 - c2)
	} else {
		s += 32 * k3
	}
	fmt.Println(s)
}
