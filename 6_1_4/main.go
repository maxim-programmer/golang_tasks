package main

import "fmt"

func main() {
	var s string
	_, _ = fmt.Scan(&s)
	flag := false
	for i := range len(s) {
		if string(s[i]) == "w" {
			fmt.Println("w")
			flag = true
			break
		} else if string(s[i]) == "x" {
			fmt.Println("x")
			flag = true
			break
		}
	}
	if !flag {
		fmt.Println(-1)
	}
}
