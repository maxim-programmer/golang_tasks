package main

import "fmt"

func main() {
	var s string
	_, _ = fmt.Scan(&s)
	i := 0
	for i <= len(s)-1 {
		if string(s[i]) == "." {
			fmt.Print(0)
			i++
		} else if string(s[i])+string(s[i+1]) == "-." {
			fmt.Print(1)
			i += 2
		} else if string(s[i])+string(s[i+1]) == "--" {
			fmt.Print(2)
			i += 2
		}
	}
}
