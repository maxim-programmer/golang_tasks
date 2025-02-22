package main

import "fmt"

func main() {
	var s string
	var cnt int = 0
	var input string
	_, _ = fmt.Scan(&s)
	for {
		_, _ = fmt.Scan(&input)
		cnt++
		if input == s {
			fmt.Println(cnt)
			break
		}
	}
}
