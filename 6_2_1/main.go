package main

import "fmt"

func main() {
	var c rune
	_, _ = fmt.Scanf("%c", &c)
	if c >= 48 && c <= 57 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
