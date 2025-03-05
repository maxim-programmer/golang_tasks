package main

import "fmt"

func main() {
	var s string
	_, _ = fmt.Scan(&s)
	s_copy := ""
	for i := range len(s) {
		s_copy = string(s[i]) + s_copy
	}
	fmt.Println(s_copy)
}
