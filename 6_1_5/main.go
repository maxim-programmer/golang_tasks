package main

import "fmt"

func main() {
	var s string
	_, _ = fmt.Scan(&s)
	palindrom := ""
	for i := range len(s) {
		palindrom = string(s[i]) + palindrom
	}
	if palindrom == s {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
