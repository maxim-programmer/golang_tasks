package main

import "fmt"

func main() {
	var s string
	_, _ = fmt.Scan(&s)
	cnt := 0
	for i := range len(s) {
		switch string(s[i]) {
		case "a":
			cnt++
		case "e":
			cnt++
		case "i":
			cnt++
		case "o":
			cnt++
		case "u":
			cnt++
		}
	}
	fmt.Println(cnt)
}
