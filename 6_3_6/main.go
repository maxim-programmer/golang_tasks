package main

import "fmt"

func main() {
	var s string
	_, _ = fmt.Scan(&s)
	for i := range len(s) {
		if string(s[i]) == "e" {
			fmt.Print("i")
		} else {
			fmt.Print(string(s[i]))
		}
	}
}
