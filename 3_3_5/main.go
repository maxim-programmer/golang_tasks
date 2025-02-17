package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var s string
	_, _ = fmt.Scan(&a, &b, &s)
	if s == "+" {
		fmt.Println(a + b)
	} else {
		if s == "-" {
			fmt.Println(a - b)
		} else {
			if s == "*" {
				fmt.Println(a * b)
			} else {
				if s == "/" {
					if b != 0 {
						fmt.Println(float64(a) / float64(b))
					} else {
						fmt.Println("На ноль делить нельзя!")
					}
				} else {
					fmt.Println("Неверная операция")
				}
			}
		}
	}
}
