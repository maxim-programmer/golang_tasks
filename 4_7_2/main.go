package main

import "fmt"

func main() {
	var (
		n   int
		cnt int = 0
		j   int
	)
	_, _ = fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		j = i
		for j > 0 {
			if j%10 == 7 {
				cnt++
			}
			j = j / 10
		}
	}
	fmt.Println(cnt)
}
