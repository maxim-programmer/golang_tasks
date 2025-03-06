package main

import "fmt"

func main() {
	var s string
	var k int
	_, _ = fmt.Scan(&s, &k)
	for i := range len(s) {
		if i == k-1 {
			fmt.Println(string(s[i]))
		} else if k > len(s) {
			fmt.Println("NO")
			break
		}
	}
}
