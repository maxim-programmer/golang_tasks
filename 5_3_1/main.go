package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	num := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&num[i])
	}
	flag := false
	for _, number := range num {
		cnt := 0
		for _, number2 := range num {
			if number == number2 {
				cnt++
			}
		}
		if cnt > 1 {
			flag = true
		}
	}
	if flag {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
