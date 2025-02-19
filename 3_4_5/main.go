package main

import "fmt"

func main() {
	var k int
	var m int
	_, _ = fmt.Scan(&k, &m)
	if k%m == 0 {
		fmt.Println(k / m)
	} else {
		fmt.Println(k/m + 1)
	}
}
