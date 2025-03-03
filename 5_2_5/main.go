package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	num := make([]int, n)
	num_fake := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&num[i])
		num_fake[n-1-i] = num[i]
	}
	flag := true
	for j := 0; j < n; j++ {
		if num[j] != num_fake[j] {
			flag = false
		}
	}
	if flag {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
