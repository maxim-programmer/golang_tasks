package main

import "fmt"

func main() {
	var cnt int
	for i := 100; i <= 999; i++ {
		if (i/100)+(i/10%10)+(i%10) == 8 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
