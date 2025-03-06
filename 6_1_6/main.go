package main

import "fmt"

func main() {
	var s string
	_, _ = fmt.Scan(&s)
	for i := range len(s) {
		cnt := 0
		for j := range len(s) {
			if string(s[i]) == string(s[j]) {
				cnt++
			}
		}
		if cnt == 2 {
			fmt.Println(string(s[i]))
			break
		}
	}
}
