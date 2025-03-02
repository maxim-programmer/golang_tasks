package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	num := make([]int, n)
	flag := false
	for i := 0; i < n; i++ {
		fmt.Scan(&num[i])
	}
	for j := 1; j < n; j++ {
		if num[j]*num[j-1] > 0 {
			flag = true
		}
	}
	if flag == true {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
