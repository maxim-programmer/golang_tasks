package main

import "fmt"

func main() {
	var s1, s2 string
	_, _ = fmt.Scan(&s1, &s2)
	fmt.Println(count_b(s1) + count_b(s2))
}

func count_b(s string) int {
	cnt := 0
	for i := range len(s) {
		if string(s[i]) == "b" {
			cnt++
		}
	}
	return cnt
}
