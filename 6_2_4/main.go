package main

import "fmt"

func main() {
	var c rune
	_, _ = fmt.Scanf("%c", &c)
	if 'a' <= c && c <= 'z' {
		fmt.Println(string(c - 32))
	} else {
		fmt.Println(string(c + 32))
	}
}
