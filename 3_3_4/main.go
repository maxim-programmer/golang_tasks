package main

import "fmt"

func main() {
	var x float64
	var y float64
	_, _ = fmt.Scan(&x, &y)
	if x > 0 {
		if y > 0 {
			fmt.Println(1)
		} else {
			fmt.Println(4)
		}
	} else {
		if y > 0 {
			fmt.Println(2)
		} else {
			fmt.Println(3)
		}
	}
}
