package main

import "fmt"

func main() {
	var a rune
	var b rune
	_, _ = fmt.Scanf("%c\n%c", &a, &b)
	if a < b {
		for i := a; i <= b; i++ {
			fmt.Print(string(i), " ")
		}
	} else {
		for i := b; i <= a; i++ {
			fmt.Print(string(i), " ")
		}
	}
}
