package main

import "fmt"

func main() {
	var num int
	_, _ = fmt.Scan(&num)
	digit1 := num / 100000
	digit2 := num / 10000 % 10
	digit3 := num / 1000 % 10
	digit4 := num / 100 % 10
	digit5 := num / 10 % 10
	digit6 := num % 10
	if (digit1 + digit2 + digit3) == (digit4 + digit5 + digit6) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
