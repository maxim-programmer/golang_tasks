package main

import "fmt"

func main() {
	var c rune
	_, _ = fmt.Scanf("%c", &c)
	for i := 'a'; i <= c; i++ {
		fmt.Print(string(i), " ")
	}
}
