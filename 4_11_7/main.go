package main

import "fmt"

func main() {
	var (
		a    int
		b    int
		max  int
		flag bool = false
	)
	_, _ = fmt.Scan(&a, &b)
	for i := a; i <= b; i++ {
		if i%7 == 0 {
			max = i
			flag = true
		}
	}
	if flag == true {
		fmt.Println(max)
	} else {
		fmt.Println("NO")
	}
}
