package main

import "fmt"

func main() {
	var a int
	var b int
	var c int
	_, _ = fmt.Scan(&a, &b, &c)
	if a == b {
		if b == c {
			fmt.Println(3)
		} else {
			fmt.Println(2)
		}
	} else {
		if b == c {
			if a == b {
				fmt.Println(3)
			} else {
				fmt.Println(2)
			}
		} else {
			if a == c {
				if a == b {
					fmt.Println(3)
				} else {
					fmt.Println(2)
				}
			} else {
				fmt.Println(0)
			}
		}
	}
}
