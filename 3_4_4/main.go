package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	if (n/100) < (n/10%10) && (n/10%10) < (n%10) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
