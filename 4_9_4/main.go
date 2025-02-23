package main

import "fmt"

func main() {
	for i := 10; i <= 99; i++ {
		if 2*(i/10)*(i%10) == i {
			fmt.Print(i)
		}
	}
}
