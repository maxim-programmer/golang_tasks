package main

import "fmt"

func main() {
	var y int
	_, _ = fmt.Scan(&y)
	for true {
		y++
		if (y/1000 != y/100%10) && (y/1000 != y/10%10) && (y/1000 != y%10) && (y/100%10 != y/10%10) && (y/100%10 != y%10) && (y/10%10 != y%10) {
			fmt.Println(y)
			break
		}
	}
}
