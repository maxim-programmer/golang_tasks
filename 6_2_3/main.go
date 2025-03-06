package main

import "fmt"

func main() {
	var c rune
	_, _ = fmt.Scanf("%c", &c)
	for i := c; i <= 'z'; i++ {
		fmt.Print(string(i), " ")
	}
}
