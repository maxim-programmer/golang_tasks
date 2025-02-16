package main

import "fmt"

func main() {
	var x1 int
	var y1 int
	var x2 int
	var y2 int
	_, _ = fmt.Scan(&x1, &y1, &x2, &y2)
	if x1 == x2 || y1 == y2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
