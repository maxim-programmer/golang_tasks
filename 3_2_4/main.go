package main

import "fmt"

func main() {
	var x int
	_, _ = fmt.Scan(&x)
	digit1 := x / 100
	digit2 := x / 10 % 10
	digit3 := x % 10
	if digit1 != digit2 && digit2 != digit3 && digit1 != digit3 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
